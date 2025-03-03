package qwen

import (
	"testing"

	"github.com/STX5/llm-translate/util"
)

var apikey = util.ApiKey

func TestChat(t *testing.T) {
	c := NewClient(apikey, true)

	resp, err := c.Chat(&util.ChatRequest{
		Model: util.MODEL_QWEN_MAX,
		Input: util.Input{
			Prompt: "你好",
		},
		Parameters: util.Parameters{
			ResultFormat: util.RESULT_FORMAT_MESSAGE,
		},
	})
	if err != nil {
		t.Errorf("err-> %v", err)
		return
	}
	t.Logf("ali: resp-> %#v", resp)
	t.Logf("ali: AI回复-> %s", resp.Output.Choices[0].Message.Content)
}

func TestChatStream(t *testing.T) {
	c := NewClient(apikey, true)

	streamReader, err := c.ChatStream(&util.ChatRequest{
		Model: util.MODEL_QWEN_MAX,
		Input: util.Input{
			Prompt: "你好",
		},
		Parameters: util.Parameters{
			ResultFormat:      util.RESULT_FORMAT_MESSAGE,
			IncrementalOutput: true,
		},
	})
	if err != nil {
		t.Errorf("err-> %v", err)
		return
	}

	defer streamReader.Close()
	for {
		line, err := streamReader.Receive()
		if err != nil {
			t.Errorf("err-> %v", err)
			break
		}
		if streamReader.IsFinish() {
			t.Logf("read finish...")
			break
		}
		if streamReader.IsMaxEmptyLimit() {
			t.Errorf("read empty limit...")
			break
		}

		if len(line) == 0 {
			continue
		}
		t.Logf("ali: resp line-> %s, len-> %d", line, len(line))
	}
}

func TestChatStreamFormat(t *testing.T) {
	c := NewClient(apikey, true)

	streamReader, err := c.ChatStream(&util.ChatRequest{
		Model: util.MODEL_QWEN_MAX,
		Input: util.Input{
			Prompt: "你是谁",
		},
		Parameters: util.Parameters{
			ResultFormat:      util.RESULT_FORMAT_MESSAGE,
			IncrementalOutput: true,
		},
	})
	if err != nil {
		t.Errorf("err-> %v", err)
		return
	}

	defer streamReader.Close()
	for {
		data, err := streamReader.ReceiveFormat()
		if err != nil {
			t.Errorf("err-> %v", err)
			break
		}
		if streamReader.IsFinish() {
			t.Logf("read finish...")
			break
		}
		if streamReader.IsMaxEmptyLimit() {
			t.Errorf("read empty limit...")
			break
		}

		if data == nil {
			continue
		}
		t.Logf("ali: resp data-> %#v", data)
	}
}
