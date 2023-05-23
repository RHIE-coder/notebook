# Webpack

번들러: 서로 연관이 있는 모듈 간의 관계를 해석하여 정적인 자원으로 변환해주는 변환 도구

애플리케이션 동작과 관련된 여러 개의 파일(HTML, CSS, JS, IMG ...)들을 1개의 자바스크립트 파일 안에 다 넣어버리고 해당 자바스크립트 파일만 로딩해도 웹 앱이 돌아가게 하자.

## [주요 속성]

### - entry

웹팩으로 빌드(변환)할 대상 파일을 지정하는 속성. entry로 지정한 파일의 내용에는 전체 애플리케이션 로직과 필요 라이브러리를 로딩하는 로직이 들어감

### - output

웹팩으로 빌드한 결과물의 위치와 파일 이름 등 세부 옵션을 설정하는 속성

### - module

로더를 지정해줌

### - loader

웹팩으로 빌드할 때 HTML, CSS, PNG 파일 등을 JS로 변환하기 위해 필요한 설정을 정의하는 속성

### - plugin

웹팩으로 빌드하고 나온 결과물에 대해 추가 기능을 제공하는 속성. 결과물의 사이즈를 줄이거나 결과물을 기타 CSS, HTML 파일로 분리하는 기능

### - resolve

웹팩으로 빌드할 때 해당 파일이 어떻게 해석되는지 정의하는 속성. 특정 라이브러리를 로딩할 때 버전은 어떤 걸로 하고, 파일 경로는 어디로 지정하는지 등을 정의

### - performance

웹팩으로 빌드한 파일의 크기가 250KB를 넘으면 경고 메시지를 표시할지 결정

### - devtool

웹팩으로 빌드된 파일로 웹앱을 구동했을 때 개발자 도구에서 사용할 디버깅 방식 지정.