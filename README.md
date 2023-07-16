# go

Go! 그것은 신인가?

# Go To Song 요약 (내기준)

## TinyGo for Maker

1. 1학년 전기전자시간
2. Mcu 개쩌러!!!
3. mcu = c/c++, assembly로 개발
4. MCU IDE는 KEIL, ill, cubeIDE(STM), arm MBED 등 을 사용함
5. Tinygo는 VSC를 사용함
6. Tinygo란 mcu를 사용하기 위한 컴파일러이다. go쓰는 사람이 iot를 개발할수 있게 만들었다.
7. Tiny PlayGround도 있음
8. TinyGo는 정말 여러가지 라이브러리가 있음
9. ”brew tap tinygo-org/tools/brew install tinygo”로 mac은 install할수있음
10. 라즈베리파이도 tinygo가능
11. I2C는 필립스에서 만든 직렬버스

## Go 언어에서 JWT와 Radius로 인증체계 만들기

1. 세션과 쿠키로 인증을 한다.
2. 세션과정을 검색하기
3. 세션과 쿠키의 취약점: http에서만 가능한데, https인 경우 ssl을 통해서 할수있다.
4. JWT의 구조
    1. 웹표준으로 지정되어있어 가볍고 안전하게 사용할 수 있다.
    2. Header: header는 두가지 타입이 있음
    3. Payload: 토큰에 잠을 정보가 담겨있고 정보의 한 조각을 ‘클레임’이라고 한다
5. 매번 번거롭게 로그인하지 않고, 자동로그인 기능을 할려면 유효시간이 유지되는 토큰이 필요한데 이 토큰을 Refrsh Token이라고 한다.
6. 로그인 플로우 사진 넣기
7. Refresh Token에 유효시간을 더 체계적으로 관리하기 위해 Radius를 사용함
8. Radius엔 ttl기능이 있어서 사용한다.

## Dagger(feat. Go): 컨테이너에서 실행되는 CI/CD엔진

1. CI/CD란
    1. 데브옵스 엔지니어의 업무
    2. 지속전 통합/배포를 위한것
    3. CI: 빌드 및 테스트를 자동화 하는 소프트웨어 개발 방법론
    4. 테스트를 하면 좋은 점: 프로젝트 안정성 증가/커뮤니케이션 비용 감소/일관된 품질로 운영 가능
    5. CI를 사용하는 이유: 테스트를 통해 얻는 이점을 CI를 통해 자동화해서 qa엔지니어들과 안정성을 올릴 수 있다.
    6. CD: 배포 과정 자동화
    7. 배포 시스템이 없다면?: 배포 히스토리 아시는 분 있나요?-히스토리에 의존한 배포/제 컴퓨터에선 작동하는데요?-자유분방한 배포 환경/인간은 같은 실수를 반복한다.-휴먼에러
    8. 배포를 시스템화화면 좋은 점: 코드 기반으로 이해가능/빌드, 배포 환경 통일 가능/사람보다 안정적이고 신속한 배포 가능
2. Dagger 소개
    1. Dagger란
        1. 컨테이너에서 파이프라인을 실행하는 프로그래밍 가능한 CI/CD엔진
    2. Dagger의 좋은 점
        1. CI/CD를 플랫폼에 구애받지 않고 즉각적인 로컬 환경 테스트 가능
        2. 모든 작업은 기본적으로 캐싱이 제공됨
        3. Dagger는 파이프라인을 표준 OCI 컨테이너로 실행함
    3. Dagger를 지원하는 SDK
        1. Node JS
        2. GraphQL
        3. GO
        4. Dagger CLI
        5. Python
    4. 성숙도에 대한 보장을 할 수 없다. 메이저 버전이 안나옴
    5. Dagger를 추천하는 이유
        1. 유망하고 빠르게 성장하고 있는 CI/CD 프로젝트임

## Go로 진행하는 AI 서비스. TensorFlow for Go

1. Background
    1. 비전테스크를 굉장히 빠른 Go로 할수있을꺼 같아서
2. TensorFlow
    1. 구글에서 만든 오픈소스화 된 머신러닝 프레임워크
    2. TensorFlow 생태계
    3. TensorFlow APIs
3. TensorFlw for Go
    1. Tensorflow
    2. tfgo
4. Conclusion
    1. 된다…
    2. TensorFlow Go를 쓰려면 Operation System의 Tensorflow Library 레벨의 호환성과 모듈을 잘 맞춰준다면 가능은 하다
    3. 모델학습은 Python으로 쓰는게 제일 좋고, 서빙쪽은 c++보단 좋은거 같다.
    4. AI할꺼면 Go를 쓰지 말자!

## Concurrent Programming in Go

1. Concurrent Programming
    1. 로직을 동시에 처리를 잘 할 수 있도록 하는것
    2. Concurrency is not parallelism
        1. Golang은 Message Passing Communication방식을 사용하여 Concurrency가 된다.
        2.
2. Go native Concurrency Primitives
    1. Goroutine은 함수를 실행 시킨다
        1. 경량쓰레드이다.
        2. func를 동시에 실행할수있고
        3. 엄청 가볍고
        4. 사용하기 쉽다.
    2. Channel은 Goroutine간에 커뮤니케이션을 하는 곳
        1. MPMC이다.
        2. 쓰레드를 바로 접근해도 안전하다
    3. Select: 채널간에 선택을 하는곳
3. Tips
    1. Goroutine이 리턴이 잘 되나 체크한다
