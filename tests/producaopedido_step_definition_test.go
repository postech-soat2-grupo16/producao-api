package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/postech-soat2-grupo16/producao-api/adapters/producaopedido"
	"github.com/postech-soat2-grupo16/producao-api/entities"
)

func parameterID(pedidoID int) error {
	inputs.pedidoID = pedidoID
	return nil
}

func requestPOSTProducaoPedido() error {
	pedidoItem := producaopedido.ProducaoPedido{PedidoID: uint32(inputs.pedidoID)}
	body, err := json.Marshal(pedidoItem)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/producao_pedidos", baseURL), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func requestPUTProducaoPedidoWithStatus(status string) error {
	pedidoItem := producaopedido.ProducaoPedido{PedidoID: uint32(inputs.pedidoID), Status: status}
	body, err := json.Marshal(pedidoItem)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/producao_pedidos/%d", baseURL, inputs.pedidoID), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func statusCodeShouldBe(statusCode int) error {
	if inputs.statusCode != statusCode {
		return fmt.Errorf("expected status code %d, got %d", statusCode, inputs.statusCode)
	}
	return nil
}

func requestGETProducaoPedidoById() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/producao_pedidos/%d", baseURL, inputs.pedidoID), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	inputs.body = res.Body
	return nil
}

func parameterStatus(status string) error {
	inputs.status = status
	return nil
}

func requestGETProducaoPedidoBy(arg1 string) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/producao_pedidos?status=%s", baseURL, inputs.status), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func requestGETProducaoPedido() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/producao_pedidos", baseURL), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func responseShouldMatchJson(arg1 *godog.DocString) error {
	var expected entities.ProducaoPedido
	if err := json.Unmarshal([]byte(arg1.Content), &expected); err != nil {
		return err
	}
	var producaoPedido entities.ProducaoPedido
	err := json.NewDecoder(inputs.body).Decode(&producaoPedido)
	if err != nil {
		return err
	}

	if expected.Status != producaoPedido.Status || expected.PedidoID != producaoPedido.PedidoID {
		return fmt.Errorf("expected %v, got %v", expected, producaoPedido)
	}
	return nil
}

func requestGETHealthcheck() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/producao_pedidos/healthcheck", baseURL), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	inputs.statusCode = res.StatusCode
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Parameter ID: (\d+)$`, parameterID)
	ctx.Step(`^request POST \/producaoPedido$`, requestPOSTProducaoPedido)
	ctx.Step(`^request PUT \/producaoPedido with status "([^"]*)"$`, requestPUTProducaoPedidoWithStatus)
	ctx.Step(`^statusCode should be (\d+)$`, statusCodeShouldBe)
	ctx.Step(`^request GET \/producaoPedido by id$`, requestGETProducaoPedidoById)
	ctx.Step(`^Parameter status "([^"]*)"$`, parameterStatus)
	ctx.Step(`^request GET \/producaoPedido by "([^"]*)"$`, requestGETProducaoPedidoBy)
	ctx.Step(`^request GET \/producaoPedido$`, requestGETProducaoPedido)
	ctx.Step(`^response should match json:$`, responseShouldMatchJson)
	ctx.Step(`^request GET \/healthcheck$`, requestGETHealthcheck)
}

var inputs Input

type Input struct {
	pedidoID   int
	statusCode int
	status     string
	body       io.ReadCloser
}
