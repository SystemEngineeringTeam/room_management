#include <stdlib.h>
#include <nfc/nfc.h>
#include <stdio.h>
#include <string.h>
#include <math.h>
#include <unistd.h>
#include <curl/curl.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <curl/curl.h>

#include <wiringPi.h>
#define pin 7

// $ gcc -o quick_start_example1 quick_start_example1.c -lnfc

// #define URL "https://www.tikuwa.monster/echo/"
#define URL "http://172.16.6.4:8081/card/"
#define SLEEP_TIME (2)

// #define SENDUID_A "{\"uid\":\""

void sendUid(char uid[])
{

  CURL *curl;
  CURLcode res;
 
  /* In windows, this will init the winsock stuff */ 
  curl_global_init(CURL_GLOBAL_ALL);
 
  /* get a curl handle */ 
  curl = curl_easy_init();
  if(curl) {
    /* First set the URL that is about to receive our POST. This URL can
       just as well be a https:// URL if that is what should receive the
       data. */ 
    curl_easy_setopt(curl, CURLOPT_URL, URL);
    /* Now specify the POST data */ 

//データ送信形式作成

    char sendUid[255];
    sprintf(sendUid,"{\"uid\":\"%s\"}",uid);
    // strcat(sendUid,uid);
    // strcat(sendUid,"\"}");


    curl_easy_setopt(curl, CURLOPT_POSTFIELDS, sendUid);
 
    /* Perform the request, res will get the return code */ 
    res = curl_easy_perform(curl);
    /* Check for errors */ 
    if(res != CURLE_OK)
      fprintf(stderr, "curl_easy_perform() failed: %s\n",
              curl_easy_strerror(res));
 
    /* always cleanup */ 
    curl_easy_cleanup(curl);
  }
  curl_global_cleanup();
}

static void print_hex(const uint8_t *pbtData, const size_t szBytes)
{
  size_t szPos;
  char hoge[32];
  char uid[szBytes*5];

  for (int i = 0; i < szBytes*5; i++)
  {
    uid[i] = '\0';
  }
  int counter=0;
  int i=0;

  for (szPos = 0; szPos < szBytes; szPos++) {
    

     sprintf(hoge,"%03d",pbtData[szPos]);

    for(i = 0;i < 3;i++){
      uid[counter]=hoge[i];
      counter++;
    }
    uid[counter]=' ';
    counter++;

    //  strcat(uid, hoge);
    //  if(szPos != szBytes-1){
    //    strcat(uid," ");
    //  }


    printf("%03d ", pbtData[szPos]);
  }

//最後の空白を抜く
  counter--;
  uid[counter]='\0';

  printf("\n");
  printf("%s\n",uid);
  sendUid(uid);
}

int main(int argc, const char *argv[])
{
  int sleepTime = SLEEP_TIME;
  struct stat file_info;
  curl_off_t speed_upload, total_time;

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
  };

  //送信用関数 8ko
  char sendUid[255] = "{\"uid\":\"";

  wiringPiSetup () ;
  pinMode (pin, OUTPUT) ;

  printf("searching...\n");
  while (1)
  {
    if (nfc_initiator_poll_target(pnd, nmMifare, 2, 255, 2, &nt) > 0) {
        printf("以下UIDだよ\n");
      printf("UID:");
      //情報のポインタと数を送信してるよ
      print_hex(nt.nti.nai.abtUid, nt.nti.nai.szUidLen);
//      sleep(sleepTime);

      for (int i = 0;i<2;i++)
      {
        digitalWrite (pin, HIGH) ; sleep (500) ;
        digitalWrite (pin,  LOW) ; sleep (500) ;
      }

      printf("searching...\n");
    }else{
      continue;
    }
    //遅延処理
  }

  nfc_close(pnd);
  nfc_exit(context);
  exit(EXIT_SUCCESS);
}

