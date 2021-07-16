# minefit_auth_demo_07-21

this project is using for demo auth system @M07-Y21


input example ------------------   for test

query {
	links{
    username,
    clinicid,
    user{
      username
    }
    
  }
}

mutation {
  createSignout(input: {username: "Parin", clinicid:"some where",signouttime:""}){
    username,
    user{
      username
    }
    clinicid
  }
}


------------------------> for using
##### timestamp when someone who logout (data from front)
mutation create{
  createSignout(input: {username: "parin", clinicid: "somewhere", signouttime: "12345678"}){
    username,
    clinicid,
    signouttime,
  }
}

>> mutation create{createSignout(input: {username: \"${..}\", clinicid: "somewhere", signouttime: \"${..}\"}){username,clinicid,signouttime,}}

##### find logout logging
query {
  links {
    username
    clinicid
    signouttime
  }
}
>> http://localhost:8080/query?query={links{username,clinicid,signouttime}}


##### create user
mutation {
  createUser(input: {
    username: "user1", 
    password: "123",
    createtime:"",
    clinicid:"A123",
    roleid:1,
    accountstatus:1,
    sessionlifetime:1000
  
  })
}
>> http://localhost:8080/query?mutation={ createUser(input: {username: \"user1\", password: \"123\", createtime:\"\",  clinicid:\"A123\",roleid:1,accountstatus:1,sessionlifetime:1000 })}

##### login
mutation {
  login(input: {
    username: "user1", 
    password: "123",
    attempt:1,
    signintime:""
  })
}
>>"mutation {login(input: {username: \"${email_data}\", password: \"${password_data}\",attempt:1,signintime:\"\"})}");

##### refreshToken example!!!!
mutation {
  refreshToken(input: {  token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjU0MTI1NzIsInVzZXJuYW1lIjoidXNlcjEifQ.ks-5hmVAVwsQtgdq4uEVjh9CFfCujBJekLdPV7w9Vao"
    
  })
}


curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"mutation { createUser(input: {username: \"user1\", password: \"123\",  createtime:\"\",  clinicid:\"A123\",  roleid:1,   accountstatus:1,    sessionlifetime:1000  })}"}' 


curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"mutation {\n  createUser(input: {\n    username: \"test@mail.com\", \n    password: \"1234567890\",\n    createtime:\"\",\n    clinicid:\"A123\",\n    roleid:1,\n    accountstatus:1,\n    sessionlifetime:1000\n  \n  })\n}\n\n"}' --compressed



##### for demo test purposes (no database)

mutation {
  login(input: {
    username: "parin@mail.com", 
    password: "123456789",
    attempt:1,
    signintime:""
  }){
    Status
    Sessionid
    Userid
    Clinicid
    Roleid
    Accountstatus
    Sessionlifetime
  }
}