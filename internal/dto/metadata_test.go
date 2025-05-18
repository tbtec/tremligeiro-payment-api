package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetaDataMarshalUnmarshal(t *testing.T) {
	paymentId := "123"
	webhookUrl := "https://webhook.url"
	meta := MetaData{
		PaymentId:         &paymentId,
		PaymentWebHookUrl: &webhookUrl,
	}
	content := MetaDataContent{Content: meta}

	data, err := json.Marshal(content)
	assert.NoError(t, err)

	var unmarshaled MetaDataContent
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	assert.NotNil(t, unmarshaled.Content.PaymentId)
	assert.NotNil(t, unmarshaled.Content.PaymentWebHookUrl)
	assert.Equal(t, paymentId, *unmarshaled.Content.PaymentId)
	assert.Equal(t, webhookUrl, *unmarshaled.Content.PaymentWebHookUrl)
}

func TestMetaDataNilFields(t *testing.T) {
	meta := MetaData{}
	content := MetaDataContent{Content: meta}

	data, err := json.Marshal(content)
	assert.NoError(t, err)

	var unmarshaled MetaDataContent
	err = json.Unmarshal(data, &unmarshaled)
	assert.NoError(t, err)
	assert.Nil(t, unmarshaled.Content.PaymentId)
	assert.Nil(t, unmarshaled.Content.PaymentWebHookUrl)
}
