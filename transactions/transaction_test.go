package transactions

import "testing"

func TestGetJSONWithOnlyParams(t *testing.T) {
	tv := DuckartParams{Amount: 1000000000000000000, From: "Tom", MaxFee: 12345, Nonce: 1, To: "Tim", Type: 127}
	t1 := DuckatTransaction{Params: &tv}

	expected_value := "{\"params\":{\"amount\":1000000000000000000,\"from\":\"Tom\",\"max_fee\":12345,\"nonce\":1,\"to\":\"Tim\",\"type\":127}}"
	actual_value := t1.GetJSON()

	if actual_value != expected_value {
		t.Errorf("expected %v but got %v", expected_value, actual_value)

	}
}
