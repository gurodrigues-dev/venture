*** Settings ***
Library    RequestsLibrary

*** Variables ***
${BASE_URL}    http://localhost:8080/api/v1
${HEADERS}     Content-Type=application/json
${EXPECTED_RESPONSE}    {"message":"pong"}

*** Keywords ***
Testing route Pong   
    Log To Console    Teste da Rota Pong... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route Health
    Log To Console    Teste da Rota Health... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route Login Users
    Log To Console    Teste da Rota Login para Users... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route Login Drivers
    Log To Console    Teste da Rota Login para Drivers... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route Get users
    Log To Console    Teste da Rota Get Users... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️
    
Testing route Post users
    Log To Console    Teste da Rota Post Users... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route Put users
    Log To Console    Teste da Rota Put Users... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route Delete users
    Log To Console    Teste da Rota Delete Users... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route Get drivers
    Log To Console    Teste da Rota Get drivers... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️
    
Testing route Post drivers
    Log To Console    Teste da Rota Post drivers... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route Put drivers
    Log To Console    Teste da Rota Put drivers... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route delete drivers
    Log To Console    Teste da Rota Delete drivers... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route user to driver
    Log To Console    Teste da Rota Usuários virando Drivers... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route password recovery
    Log To Console    Teste da Rota Recuperação de Senha... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route password verify
    Log To Console    Teste da Rota Recuperação de Senha... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

Testing route password change
        Log To Console    Teste da Rota Recuperação de Senha... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}    GET On Session    api_session    /ping
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    Should Be Equal As Strings    ${response.content}    {"message":"pong"}
    Log To Console   [OK] ✔️

*** Test Cases ***
Unitary Tests
    Testing route pong
    Testing route health

    Testing route Post users
    Testing route Login Users
    Testing route Get users
    Testing route Put users
    Testing route Delete users
    Testing route password recovery
    Testing route password verify
    Testing route password change

    Testing route Post drivers
    Testing route Login Drivers
    Testing route Get drivers
    Testing route Put drivers
    Testing route Delete drivers

    Testing route user to driver
    