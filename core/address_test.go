package core

import (
	"crypto/rsa"
	"math/big"
	"testing"
)

/*
{
	"state": {
		"address": "0x236d3dd714763154fb40a223588501fd02000200",
		"value": 0,
		"code": null,
		"data": null,
		"nonce": 0
	},
	"private_key": {
		"N": 116406071470386766331899700083503647390791063720795371113019376328493291782063189055057503854190979083526546854461552541078574769541891907776008090005630958764053549493899702227537714232482524424991649990549054020277588865012308749074304734527683005317741142699097108044848797455182949464572553404571539691653,
		"E": 65537,
		"D": 60862878145235102844635474969274394334388919075587442003123761375591667854863317456713008332204832221128837977249180758542478951632245420325200927905350408048714772929483599113355329003459548999518674376916443837918933297608434964331275672791218707521943040754870597966182662029946432385546964362491082705153,
		"Primes": [
			10164045354392978158556024837533195250057636899646752306346023543763742840616738026538342414555844812591542003187056340470307163148224592498711327398381413,
			11452730424906572432742546315183506953851932601197705199789887818184635953795752527137768234058681310589855746977110912076594800886302386188953893820664481
		],
		"Precomputed": {
			"Dp": 3499730647836061234508373231667212170419772560193915079191973500291020659190339965315230097912128935424885741549340271587850396624850636347802307308419289,
			"Dq": 3224526325592811305506285985134596797419409804496697382341623968462744749064486101450266253488117972183713295439547897365416898331538697681593256462439553,
			"Qinv": 808348842325348247110680490961273605770102610864303729411444821011318134672535046469196171684533796148431232770145511677730295490061078858933874249480786,
			"CRTValues": []
		}
	},
	"public_key": {
		"N": 116406071470386766331899700083503647390791063720795371113019376328493291782063189055057503854190979083526546854461552541078574769541891907776008090005630958764053549493899702227537714232482524424991649990549054020277588865012308749074304734527683005317741142699097108044848797455182949464572553404571539691653,
		"E": 65537
	},
	"address": "0x236d3dd714763154fb40a223588501fd02000200"
}
*/

const publicKeyN string = "116406071470386766331899700083503647390791063720795371113019376328493291782063189055057503854190979083526546854461552541078574769541891907776008090005630958764053549493899702227537714232482524424991649990549054020277588865012308749074304734527683005317741142699097108044848797455182949464572553404571539691653"
const publicKeyE int = 65537
const publicKeyHex string = "24ff81030101095075626c69634b657901ff8200010201014e01ff840001014501040000000aff83050102ff86000000ff8cff8201ff8102a5c48b25ca685ff9f2da8201cfe93a7eca2f938071593098a65d1bf73d1be5a7d4efa13b20c8d6af67a1175ca79b3224056a2f27f36624ee61a0278455b72284ab4c01ca876b7e05a06504849f8a40d757c919f287389a3b36d6b89d3ee073df742a6237c5047f219cfc13db9059fdb89c28236d3dd714763154fb40a223588501fd02000200"
const publicKeyAddress string = "0x236d3dd714763154fb40a223588501fd02000200"

type CoreTestingStruct struct {
	publicKey      rsa.PublicKey
	hexKey         string
	accountAddress string
}

func TestingSetup() CoreTestingStruct {
	nValue, _ := new(big.Int).SetString(publicKeyN, 10)
	return CoreTestingStruct{
		publicKey:      rsa.PublicKey{N: nValue, E: publicKeyE},
		hexKey:         publicKeyHex,
		accountAddress: publicKeyAddress,
	}
}

func TestPublicKeyToHex(t *testing.T) {
	testingStruct := TestingSetup()

	actual_value := PublicKeyToHex(testingStruct.publicKey)

	if testingStruct.hexKey != actual_value {
		t.Errorf("expected %v but got %v", testingStruct.hexKey, actual_value)
	}
}

func TestHexToPublicKey(t *testing.T) {
	testingStruct := TestingSetup()

	actual_value := HexToPublicKey(testingStruct.hexKey)

	if actual_value.Equal(testingStruct.publicKey) {
		t.Errorf("expected %v but got %v", testingStruct.publicKey, actual_value)
	}
}

func TestGetAddress(t *testing.T) {
	testingStruct := TestingSetup()

	actual_value := GetAddress(testingStruct.hexKey)

	if testingStruct.accountAddress != actual_value {
		t.Errorf("expected %v but got %v", testingStruct.accountAddress, actual_value)
	}
}