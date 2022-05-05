# go-backend-with-cleancode-and-tdd
[![Maintainability](https://api.codeclimate.com/v1/badges/26b230fc84f3d4f9dac0/maintainability)](https://codeclimate.com/github/hyun06000/go-backend-with-cleancode-and-tdd/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/26b230fc84f3d4f9dac0/test_coverage)](https://codeclimate.com/github/hyun06000/go-backend-with-cleancode-and-tdd/test_coverage)  
  
코드에 관한 상세한 명세는 [위키](https://github.com/hyun06000/go-backend-with-cleancode-and-tdd/wiki)를 참고하여 주시기바랍니다.


## Summary
이 저장소는 GO 언어를 통해 API서버를 빌드하고 배포하는 것을 기록합니다.  
애자일 프로젝트 정신에 입각하여 TDD와 CleanCode를 최대한 반영하기 위해 노력합니다.  
GO언어를 이용한 서버 빌드와 GO언어가 가진 특성들을 이해하며  
코드로 정리하는 것을 목적으로 두고 있습니다.  

## Code Convention
[CleanCode에 따른 convention](https://github.com/hyun06000/go-backend-with-cleancode-and-tdd/wiki/CleanCode%EC%97%90-%EB%94%B0%EB%A5%B8-convention)을 통해 CleanCode를 어떻게 리펙토링에 적용하는지 설명합니다.  
  
## Test Driven Development  
[TDD와 cleancode로 만들면서 배우면서 Golang backend](https://davi06000.tistory.com/137?category=925226) 시리즈를 포스팅하며 테스트를 어떻게 적용하는지 명시합니다.  
[Code Climate](https://codeclimate.com/github/hyun06000/go-backend-with-cleancode-and-tdd) 을 통해 Maintainability와 Test Coverage를 기록하고 관리합니다.
Code Climate는 Github Actions로 연동하여 Continuous Integration이 적용되도록 운영하고 있습니다.  
  
## Fake Database Object for In-memory Test  
DB 기능 수행과 관련한 In-memory Test를 구현하기 위해 Fake Database Object를 구현하여 처리합니다. 자세한 내용은 [여기](https://github.com/hyun06000/go-backend-with-cleancode-and-tdd/wiki/%5B-Docs-%5D-FakeDB-module)를 참고해 주시기 바랍니다.  

## MySQL CI  
Unit-test 뿐만 아니라 MySQL을 이용하여 직접 DB를 올리고 연동하는 통합테스트를 위해 [MySQL GitHub Action](https://github.com/marketplace/actions/setup-mysql#:~:text=MySQL%20v1%20release-,MySQL%20GitHub%20Action,MySQL%2C%20see%20The%20Default%20MySQL.) 를 `workflows`에 추가하여 `git push`가 진행될 때 마다 실재 DB와 연동을 테스트합니다. 연동을 위한 코드는 [여기](https://github.com/hyun06000/go-backend-with-cleancode-and-tdd/tree/main/mysqlModule) 에서 확일할 수 있습니다. Github Action을 통한 자동화된 Testlog는 [여기](https://github.com/hyun06000/go-backend-with-cleancode-and-tdd/actions)에서 확인할 수 있습니다.  


## Reference
TDD를 통해 서버를 빌드하는 과정은 [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/)를 참고하여 진행합니다.  
[TDD와 cleancode로 만들면서 배우면서 Golang backend](https://davi06000.tistory.com/137?category=925226) 시리즈를 포스팅하며 자세한 reference와 내용을 기록합니다.
  
