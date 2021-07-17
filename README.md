# minefit_auth_demo2_07-21

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
>>"mutation {login(input: {username: \"${email_data}\", password: \"${password_data}\",attempt:1,signintime:\"\"})}");

##### refreshToken example!!!!
mutation {
  refreshToken(input: {  token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjU0MTI1NzIsInVzZXJuYW1lIjoidXNlcjEifQ.ks-5hmVAVwsQtgdq4uEVjh9CFfCujBJekLdPV7w9Vao"
    
  })
}


##### for demo test purposes 

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
