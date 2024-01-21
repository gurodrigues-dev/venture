*** Settings ***
Library    RequestsLibrary
Library    Collections

*** Variables ***
${BASE_URL}    http://localhost:8080/api/v1
${HEADERS}     Content-Type=application/json
&{HEADERS-FORM}    Content-Type=application/x-www-form-urlencoded

*** Keywords ***
Login and Get Token
    [Arguments]    ${cpf}    ${password}    ${table}
    Create Session    api_session    ${BASE_URL}
    
    ${data}    Create Dictionary
        ...    cpf=${cpf}
        ...    password=${password}
        ...    table=${table}

    ${response}    POST On Session    api_session    /login    data=${data}    headers=${HEADERS-FORM}
    Log    ${response.content}

    ${json_response}    Evaluate    json.loads('''${response.content}''')    json
    Dictionary Should Contain Key    ${json_response}    token

    [Return]    ${json_response['token']}

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
    EXCEPT    message
        Log To Console    message
        Log To Console    Teste da Rota Post Users falhou ❌    
    END


Testing route Get users
    TRY
       ${token}    Login and Get Token    65246837068    123teste    users
    
        &{HEADERS-FORM-GET-USER}    Create Dictionary    Content-Type=application/x-www-form-urlencoded    Authorization=${token}
        Log To Console    Teste da Rota Get Users... 
        Sleep    1
        Create Session    api_session    ${BASE_URL}
        ${response}    GET On Session    api_session    /users/65246837068    headers=${HEADERS-FORM-GET-USER}
        Log    ${response.content}
        Should Be Equal As Strings    ${response.status_code}    200
        ${json_response}    Evaluate    json.loads('''${response.content}''')    json

        Dictionary Should Contain Key    ${json_response}    dataOfUser
        Dictionary Should Contain Key    ${json_response}    requestID

        Log To Console   [OK] ✔️ 
    EXCEPT    message
        Log To Console    message
        Log To Console    Teste da Rota Get Users falhou ❌   
    END

Testing route Delete users
    ${token}    Login and Get Token    65246837068    123teste    users

    &{HEADERS-FORM-DELETE-USER}    Create Dictionary    Content-Type=application/x-www-form-urlencoded    Authorization=${token}
    
    Log To Console    Teste da Rota Delete Users... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}   DELETE On Session    api_session    /users/65246837068    headers=${HEADERS-FORM-DELETE-USER}   
    Log    ${response.content}
    ${json_response}    Evaluate    json.loads('''${response.content}''')    json

    Dictionary Should Contain Value    ${json_response}    User deleted w/ success

    Log To Console   [OK] ✔️ 

Testing route Post drivers
    TRY
        Log To Console    Teste da Rota Post Users... 
        Sleep    1
        Create Session    api_session    ${BASE_URL}
    
        ${data}    Create Dictionary
        ...    name=User Test
        ...    cnh=59691146158
        ...    cpf=55931964002
        ...    rg=506982270
        ...    email=guuhdazueira@gmail.com
        ...    rua=Rua Jose Falchi
        ...    numero=20
        ...    cidade=sao paulo
        ...    estado=sp
        ...    cep=04921090
        ...    password=123teste


        ${response}    POST On Session    api_session    /drivers    data=${data}    headers=${HEADERS-FORM}
        Log    ${response.content}
        ${json_response}    Evaluate    json.loads('''${response.content}''')    json

        Dictionary Should Contain Key    ${json_response}    email
        Dictionary Should Contain Key    ${json_response}    requestID
        Dictionary Should Contain Key    ${json_response}    status
        Log To Console   [OK] ✔️
    EXCEPT    message
        Log To Console    message
        Log To Console    Teste da Rota Post Drivers falhou ❌    
    END

Testing route Get drivers
    TRY
       ${token}    Login and Get Token    55931964002    123teste    drivers
    
        &{HEADERS-FORM-GET-DRIVER}    Create Dictionary    Content-Type=application/x-www-form-urlencoded    Authorization=${token}
        Log To Console    Teste da Rota Get Drivers... 
        Sleep    1
        Create Session    api_session    ${BASE_URL}
        ${response}    GET On Session    api_session    /drivers/55931964002    headers=${HEADERS-FORM-GET-DRIVER}
        Log    ${response.content}
        Should Be Equal As Strings    ${response.status_code}    200
        ${json_response}    Evaluate    json.loads('''${response.content}''')    json

        Dictionary Should Contain Key    ${json_response}    dataOfDriver
        Dictionary Should Contain Key    ${json_response}    requestID

        Log To Console   [OK] ✔️ 
    EXCEPT    message
        Log To Console    message
        Log To Console    Teste da Rota Get Users falhou ❌   
    END

Testing route Delete drivers
    ${token}    Login and Get Token    55931964002     123teste    drivers

    &{HEADERS-FORM-DELETE-DRIVER}    Create Dictionary    Content-Type=application/x-www-form-urlencoded    Authorization=${token}
    
    Log To Console    Teste da Rota Delete Drivers... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}   DELETE On Session    api_session    /drivers/55931964002     headers=${HEADERS-FORM-DELETE-DRIVER}   
    Log    ${response.content}
    ${json_response}    Evaluate    json.loads('''${response.content}''')    json

    Dictionary Should Contain Value    ${json_response}    User deleted w/ success

    Log To Console   [OK] ✔️ 

Testing route user to driver
    ${token}    Login and Get Token    65246837068    123teste    users
    ${data}    Create Dictionary
        ...    cnh=38908526120
        ...    password=123teste
    &{HEADERS-FORM-POST-USER-TO-DRIVER}    Create Dictionary    Content-Type=application/x-www-form-urlencoded    Authorization=${token}
        Log To Console    Teste da Rota User virando Driver... 
        Sleep    1
        Create Session    api_session    ${BASE_URL}
    ${response}    Post On Session    api_session    /users/drivers/65246837068    data=${data}    headers=${HEADERS-FORM-POST-USER-TO-DRIVER}
    Log    ${response.content}
    ${json_response}    Evaluate    json.loads('''${response.content}''')    json
    
    Dictionary Should Contain Key    ${json_response}    s3bucketurl

    Log To Console   [OK] ✔️

Post User-Driver
    TRY
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
    EXCEPT    message
        Log To Console    message
        Log To Console    Teste da Rota Post Users falhou ❌    
    END
    
Delete User-Driver
    ${token}    Login and Get Token    65246837068    123teste    users

    &{HEADERS-FORM-DELETE-USER}    Create Dictionary    Content-Type=application/x-www-form-urlencoded    Authorization=${token}
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}   DELETE On Session    api_session    /users/65246837068    headers=${HEADERS-FORM-DELETE-USER}   
    Log    ${response.content}
    ${json_response}    Evaluate    json.loads('''${response.content}''')    json

    Dictionary Should Contain Value    ${json_response}    User deleted w/ success

    &{HEADERS-FORM-DELETE-DRIVER}    Create Dictionary    Content-Type=application/x-www-form-urlencoded    Authorization=${token}
    
    Log To Console    Teste da Rota Delete Drivers... 
    Sleep    1
    Create Session    api_session    ${BASE_URL}
    ${response}   DELETE On Session    api_session    /drivers/65246837068     headers=${HEADERS-FORM-DELETE-DRIVER}   
    Log    ${response.content}
    ${json_response}    Evaluate    json.loads('''${response.content}''')    json

    Dictionary Should Contain Value    ${json_response}    User deleted w/ success
    

*** Test Cases ***
Unitary Tests
    # server
    Testing route pong
    Testing route health

    # users
    Testing route Post users
    Testing route Get users
    Testing route Delete users

    # drivers
    Testing route Post drivers
    Testing route Get drivers
    Testing route Delete drivers

    # user to driver
    Post User-Driver
    Testing route user to driver
    Delete User-Driver