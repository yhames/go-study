# Websocket 채팅서버

## 트러블슈팅

### 윈도우에서 `confluent-kafka-go` 패키지 참조 오류

- 주요 원인: `confluent-kafka-go`는 C 라이브러리(`librdkafka`)에 의존하는 CGO 패키지인데, Windows 환경에서 다음이 문제가 발생
    - CGO가 비활성화되어 있음 (CGO_ENABLED=0)
    - C 컴파일러 부재
    - librdkafka 라이브러리 부재

- 해결 방법:
    - MSYS2를 설치하여 MinGW-w64 툴체인 제공
    - GCC 컴파일러 설치
    - librdkafka 라이브러리 설치
    - CGO 활성화 (CGO_ENABLED=1)
    - MinGW-w64 bin 디렉토리를 PATH에 추가

- 실행명령어:
  ```bash
  $env:PATH = "C:\msys64\mingw64\bin; $env:PATH"; $env:CGO_ENABLED = "1"; go run .
  ```
  
- GoLand 환경변수 설정:
  ```bash
  PATH=C:\msys64\mingw64\bin;CGO_ENABLED=1;
  ```