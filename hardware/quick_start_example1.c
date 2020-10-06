#include <stdlib.h>
#include <nfc/nfc.h>
#include <stdio.h>
#include <string.h>
#include <math.h>
#include <unistd.h>

static void print_hex(const uint8_t *pbtData, const size_t szBytes)
{
  size_t  szPos;
  char uid[szBytes];

  for (szPos = 0; szPos < szBytes; szPos++) {
    uid[szPos] = (char)pbtData[szPos];

    printf("%d ", pbtData[szPos]);
  }
  printf("\n");
  printf("charで表示するよ…\n");
  printf("%s",uid);
  printf("\n");
}

int main(int argc, const char *argv[])
{
  int sleepTime = 5;

  nfc_device *pnd;
  nfc_target nt;

  nfc_context *context;

  //nfc_initはlibnfcを初期化します
  nfc_init(&context);
  if (context == NULL) {
    printf("Unable to init libnfc (malloc)\n");
    exit(EXIT_FAILURE);
  }
  printf("init OK\n");
  //nfc_openはnfc_deviceを返します
  pnd = nfc_open(context, NULL);
//ヌルエラーを返すよ
  if (pnd == NULL) {
    printf("ERROR: %s\n", "Unable to open NFC device.");
    exit(EXIT_FAILURE);
  }
  //ヌルじゃない場合はOK出力ですすむ
  printf("open OK\n");
//nfc_initiator_initはデバイスをイニシエーター初期化しますエラーの場合負の数を返します
  if (nfc_initiator_init(pnd) < 0) {
    printf("ERROR: %s\n", "nfc_initiator_init");
    exit(EXIT_FAILURE);
  }
   //ヌルじゃない場合はOK出力ですすむ
  printf("initiator init OK\n");

  const nfc_modulation nmMifare[] = {
    { .nmt = NMT_ISO14443A, .nbr = NBR_106 },
    { .nmt = NMT_ISO14443B, .nbr = NBR_106 },
  };
  while (1)
  {
   printf("searching...\n");
    if (nfc_initiator_poll_target(pnd, nmMifare, 2, 255, 2, &nt) > 0) {
        printf("以下UIDだよ\n");
      printf("UID:");
      //情報のポインタと数を送信してるよ
      print_hex(nt.nti.nai.abtUid, nt.nti.nai.szUidLen);
    } 
    //遅延処理
    sleep(sleepTime);
  }
  nfc_close(pnd);
  nfc_exit(context);
  exit(EXIT_SUCCESS);
}
