package simpleEncryptFile

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

const (
	TEMP_FOLDER string = "./temp"
)

func TestMain(t *testing.M) {
	if err := os.Mkdir(TEMP_FOLDER, 0700); err != nil {
		fmt.Println("Failed to create temp folder")
		os.Exit(1)
	}
	resu := t.Run()

	if err := os.RemoveAll(TEMP_FOLDER); err != nil {
		fmt.Println("Failed to remove temp folder")
		os.Exit(1)
	}
	os.Exit(resu)

}

func TestGenerateKeysWithoutPass(t *testing.T) {
	testData := []string{"chave1", "outrachave", "mais_uma_chave", "chave com espaços"}
	for _, tc := range testData {
		t.Logf("Testing: %s", tc)
		path := TEMP_FOLDER + "/" + tc
		err := GenerateKeys(path, "")
		if err != nil {
			t.Errorf("Failed to generate pair of Keys, %v", err)
		}
		_, err = LoadKey(path, "")
		if err != nil {
			t.Errorf("Failed to load private key, %v", err)
		}
		_, err = LoadPubKey(path + ".pub")
		if err != nil {
			t.Errorf("Failed to load public key, %v", err)
		}
	}
}

func TestGenerateKeysWithPass(t *testing.T) {
	testData := [][]string{
		{
			"chave2",
			"TNssk29YgN^%FdZ#9D45#X7SSnR99A8tmSWH*v!dyGNFCiioHskt@",
		},
		{
			"outrachave2",
			"q5b$PYnSxAofrAb3BHMyxjbu%FRABd2ERcKvfGWfi7jikmQg^xfNV",
		},
		{
			"mais_uma_chave2",
			"!&cg^ML&9Kdh6L^n^@ZXRBhd7f8ibWJ!$HSY@qnLSyjP5E*FT&zwh",
		},
		{
			"chave com espaços2",
			"2q26UBd$!E#ZG8E#hX2zE!u!peQ8E@v%o648va@FyUAAV2*KG!DBg",
		},
	}
	for _, tc := range testData {
		t.Logf("Testing: %s", tc)
		path := TEMP_FOLDER + "/" + tc[0]
		err := GenerateKeys(path, tc[1])
		if err != nil {
			t.Errorf("Failed to generate Key, %v", err)
		}
		_, err = LoadKey(path, tc[1])
		if err != nil {
			t.Errorf("Failed to load private key, %v", err)
		}
		_, err = LoadPubKey(path + ".pub")
		if err != nil {
			t.Errorf("Failed to load public key, %v", err)
		}
	}
}

func TestEncryptAndDecrypt(t *testing.T) {
	testData := []string{
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer non fringilla mi, ut vulputate elit. Morbi vulputate velit est, eu fermentum turpis mattis ut. Aenean vel vestibulum mauris. Aenean vitae iaculis ante. Aliquam vestibulum non enim at convallis. Phasellus hendrerit, ligula sit amet congue porta, orci enim aliquam mi, sed sodales ligula dolor pellentesque dolor. Vivamus at viverra nunc. Donec pharetra consectetur risus, a pellentesque risus ultricies quis. Aliquam eget dignissim leo. Donec in gravida sem, nec eleifend justo. Cras euismod blandit lorem, id euismod turpis auctor a. Sed aliquam molestie dui. Nunc id diam sit amet odio malesuada interdum. Donec laoreet nibh a condimentum pharetra. Mauris sodales nunc eu venenatis fermentum. Nullam varius enim nisi, at tempor diam pulvinar in.

Aenean faucibus metus nunc, non ullamcorper enim condimentum eget. Donec faucibus diam quis ipsum lobortis commodo. Sed egestas ligula mi, sed semper ex aliquam sit amet. Nunc semper diam sed accumsan sollicitudin. Vivamus blandit vel tellus vel imperdiet. Mauris feugiat, nisi at eleifend condimentum, leo velit tincidunt magna, et congue quam lacus a lectus. Nullam tincidunt cursus lectus lacinia pharetra. Ut dapibus lacus ut fringilla efficitur. Praesent luctus diam elementum, hendrerit lectus quis, condimentum sapien. Mauris vehicula orci sit amet libero gravida pharetra. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Aenean vitae eleifend risus. Maecenas placerat, augue vitae suscipit iaculis, lectus ante mollis nisi, eget tincidunt elit mi et mi.

Vivamus vehicula sagittis varius. Ut finibus at dui quis porttitor. Mauris hendrerit at nisi in finibus. Nulla semper viverra nisi, id mollis dolor euismod volutpat. Proin egestas condimentum ex in egestas. Integer luctus ultrices magna, quis interdum nibh lobortis vitae. Donec imperdiet metus vel dapibus iaculis. Praesent feugiat, magna ultricies consectetur aliquam, leo metus elementum dolor, lacinia consequat mi nibh id est. Cras pulvinar finibus nisi ac aliquet. Proin consectetur augue at posuere porttitor. Fusce eget dolor at enim lobortis imperdiet. Aenean eleifend nec eros sit amet congue. Aenean vitae odio metus. Phasellus iaculis bibendum ligula, sed interdum lacus pharetra id.

Integer sit amet libero eget dolor malesuada placerat. Donec sit amet sem id mi vulputate luctus. Phasellus id lacus non justo molestie pellentesque. Etiam placerat enim sed augue finibus egestas. Etiam lacus elit, elementum in viverra nec, tincidunt a nibh. Proin ligula nunc, varius pretium elit in, aliquet pulvinar lorem. Donec purus ex, lobortis ut tempor in, volutpat sit amet nunc. Maecenas vel nisl a ex mattis dapibus. Ut commodo quam vel sollicitudin vestibulum. Curabitur justo tellus, pretium sit amet dolor vel, finibus lacinia arcu.

In dignissim sapien ipsum, ac venenatis lorem euismod non. Mauris eu ligula ut mi pulvinar aliquet id id orci. Nam cursus dolor sed neque sollicitudin ultricies. Nunc accumsan erat eget convallis bibendum. Duis diam massa, finibus at mauris vitae, ornare consectetur lacus. Etiam vitae vulputate magna. Suspendisse placerat purus in ex mollis, vel tincidunt lorem tempor. Pellentesque quis tincidunt est. Quisque sit amet vulputate nisi. Duis hendrerit tempor sem eleifend ultrices. Aenean a tincidunt erat, non cursus orci. Sed elit metus, lobortis non posuere quis, rhoncus vitae diam. Sed blandit nec mauris nec pretium. Mauris porttitor est nec maximus pulvinar. Curabitur at ex sit amet lorem dapibus convallis sed eu lectus. Cras condimentum at eros vel porttitor.`,
		"a",
		"",
	}
	GenerateKeys(TEMP_FOLDER+"/qkey", "")
	key, err := LoadKey(TEMP_FOLDER+"/qkey", "")
	if err != nil {
		t.Errorf("failed to load key %v", err)
	}
	pubKey, err := LoadPubKey(TEMP_FOLDER + "/qkey.pub")
	if err != nil {
		t.Errorf("failed to load Pubkey %v", err)
	}
	for _, v := range testData {
		var reader io.Reader
		reader = strings.NewReader(v)
		out, err := EncryptData(io.NopCloser(reader), pubKey)
		if err != nil {
			t.Errorf("Failed to encrypt data")
		}
		reader = bytes.NewReader(out)
		out, err = DencryptData(io.NopCloser(reader), key)
		outString := string(out)
		if outString != v {
			t.Log("Encrypt and Decrypt generate not equal results")
			t.Failed()
		}
	}
}
