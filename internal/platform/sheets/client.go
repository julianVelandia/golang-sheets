package sheets

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julianVelandia/golang-sheets/internal/platform/number"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"

	ErrorUseCase "github.com/julianVelandia/golang-sheets/internal/cell/core/error"
	"github.com/julianVelandia/golang-sheets/internal/cell/core/query"
	"github.com/julianVelandia/golang-sheets/internal/platform/constant"
	logPlatform "github.com/julianVelandia/golang-sheets/internal/platform/log"
	"github.com/julianVelandia/golang-sheets/internal/platform/sheets/model"
)

const (
	actionEmptyResponse           string                  = "empty response"
	actionUnableToReadClient      string                  = "Unable to read client secret file"
	actionUnableRetrieve          string                  = "Unable to retrieve data from sheet"
	actionUnableToParseSecretFile string                  = "Unable to parse client secret file to config"
	errorRead                     logPlatform.LogsMessage = "error in the use case, when read repository"
	entityType                    string                  = "read_repository"
	layer                         string                  = "client_sheets_read"
)

type ClientSheets struct{}

type SpreadsheetID struct {
	SpreadsheetID string `json:"spreadsheet_id"`
}

func (c ClientSheets) getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func (c ClientSheets) tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableToReadClient,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s`,
						file,
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func (c ClientSheets) saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func (c ClientSheets) Read(ctx context.Context, credentialsPath, spreadsheetIDPath string, readRange query.GetCells) ([]model.Cell, error) {
	credentialsRead, errCredentials := os.ReadFile(credentialsPath)
	spreadsheetIDRead, errSpreadsheet := os.ReadFile(spreadsheetIDPath)

	if errCredentials != nil || errSpreadsheet != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableToReadClient,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						credentialsRead,
						spreadsheetIDRead,
						readRange.Value(),
					),
					constant.EntityType: entityType,
				}))

		err := errSpreadsheet
		if errCredentials != nil {
			err = errCredentials
		}
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}

	config, err := google.JWTConfigFromJSON(credentialsRead, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableToParseSecretFile,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						credentialsRead,
						spreadsheetIDRead,
						readRange.Value(),
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}

	client := config.Client(oauth2.NoContext)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableRetrieve,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						credentialsRead,
						spreadsheetIDRead,
						readRange.Value(),
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}

	spreadsheetID := SpreadsheetID{}
	json.Unmarshal(spreadsheetIDRead, &spreadsheetID)
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID.SpreadsheetID, readRange.Value()).Do()
	if err != nil {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionUnableRetrieve,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						credentialsRead,
						spreadsheetIDRead,
						readRange.Value(),
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	}

	if len(resp.Values) == 0 {
		message := errorRead.GetMessageWithTagParams(
			logPlatform.NewTagParams(layer, actionEmptyResponse,
				logPlatform.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						credentialsRead,
						spreadsheetIDRead,
						readRange.Value(),
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     err,
		}
	} else {

		response := make([]model.Cell, 0)
		currentRow := readRange.StartRow
		currentColumn := readRange.StartColumn
		for _, row := range resp.Values {
			for _, val := range row {
				response = append(response, model.Cell{
					CellPosition: fmt.Sprintf("%v%v", currentColumn, currentRow),
					Information:  fmt.Sprintf("%v", val),
				})
				currentRow++
			}
			currentColumn = number.UpdateChar(currentColumn)
		}
		return response, nil
	}
}
