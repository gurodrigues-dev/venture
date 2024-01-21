*** Settings ***
Library    RequestsLibrary
Library    Collections

*** Variables ***
${BASE_URL}    http://localhost:8080/api/v1
${HEADERS}     Content-Type=application/json
&{HEADERS-FORM}    Content-Type=application/x-www-form-urlencoded

*** Keywords ***
Testing route Pong   
    TRY
        Log To Console    Teste da Rota Pong... 
        Sleep    1
        Create Session    api_session    ${BASE_URL}
        ${response}    GET On Session    api_session    /ping
        Log    ${response.content}
        Should Be Equal As Strings    ${response.status_code}    200
        Should Be Equal As Strings    ${response.content}    {"message":"pong"}
        Log To Console   [OK] ✔️
        
    EXCEPT    
        Log To Console    Teste da Rota Pong falhou ❌
    END 

Testing route Health
    TRY
        Log To Console    Teste da Rota Health... 
        Sleep    1
        Create Session    api_session    ${BASE_URL}
        ${response}    GET On Session    api_session    /health
        Log    ${response.content}
        Should Be Equal As Strings    ${response.status_code}    200
        ${json_response}    Evaluate    json.loads('''${response.content}''')    json
        
        Dictionary Should Contain Key    ${json_response}    cpu
        Dictionary Should Contain Key    ${json_response}    envs
        Dictionary Should Contain Key    ${json_response}    mem
        Dictionary Should Contain Key    ${json_response}    message
        Dictionary Should Contain Key    ${json_response}    uptime

        Log To Console   [OK] ✔️
        
    EXCEPT
        
        Log To Console    Teste da Rota Health falhou ❌

    END
    

Testing route Post users
    TRY
        Log To Console    Teste da Rota Post Users... 
        Sleep    1
        Create Session    api_session    ${BASE_URL}
    
        ${data}    Create Dictionary
        ...    name=User Test
        ...    cpf=65246837068
        ...    rg=506982270
        ...    email=roderscleysonjn@gmail.com
        ...    rua=Rua Jose Falchi
        ...    numero=20
        ...    cidade=sao paulo
        ...    estado=sp
        ...    cep=04921090
        ...    complemento=apto 5
        ...    password=123teste


        ${response}    POST On Session    api_session    /users    data=${data}    headers=${HEADERS-FORM}
        Log    ${response.content}
        Should Be Equal As Strings    ${response.status_code}    201
        ${json_response}    Evaluate    json.loads('''${response.content}''')    json

        Dictionary Should Contain Key    ${json_response}    email
        Dictionary Should Contain Key    ${json_response}    requestID
        Dictionary Should Contain Key    ${json_response}    status
        Log To Console   [OK] ✔️
    EXCEPT
        Log To Console    Teste da Rota Post Users falhou ❌    
    END

Login and Get Token
    [Arguments]    ${cpf}    ${password}    ${table}
    Log To Console    Teste da Rota Login para Users... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    
    ${data}    Create Dictionary
        ...    cpf=${cpf}
        ...    password=${password}
        ...    table=${table}

    ${response}    POST On Session    api_session    /login
    Log    ${response.content}

    ${json_response}    Evaluate    json.loads('''${response.content}''')    json
    Dictionary Should Contain Key    ${json_response}    token

    [Return]    ${json_response.token}

Testing route Get users
    ${token}    Login and Get Token    65246837068    123teste    users
    ${HEADERS-FORM}    Create Dictionary    Content-Type=application/x-www-form-urlencoded    Authorization=${token}
    Log To Console    Teste da Rota Get Users... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}

    ${data}    Create Dictionary
        ...    cpf=65246837068
        ...    password=123teste
        ...    table=users

    ${response}    GET On Session    api_session    /users/65246837068
    Log    ${response.content}
    Should Be Equal As Strings    ${response.status_code}    200
    ${json_response}    Evaluate    json.loads('''${response.content}''')    json

    Dictionary Should Contain Key    ${json_response}    dataOfUser
    Dictionary Should Contain Key    ${json_response}    requestID

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

Testing route Post drivers
    Log To Console    Teste da Rota Post drivers... 
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

Testing route Get drivers
    Log To Console    Teste da Rota Get drivers... 
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
    Testing route Get users
    # Testing route Login Users
    # Testing route Put users
    # Testing route Delete users
    # Testing route password recovery
    # Testing route password verify
    # Testing route password change

    # Testing route Post drivers
    # Testing route Login Drivers
    # Testing route Get drivers
    # Testing route Put drivers
    # Testing route Delete drivers

    # Testing route user to driver
    