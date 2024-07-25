Hospital Management System
Goal
You are expected to build a hospital management system.
Tech Stack:
golang - gin (framework) for routing
Mysql/Postgres - Gorm as orm library
Requirements:
The system will accept a list of doctors. Each list element will contain the doctor’s name, address, and phone number.
The system will then accept a list of patients. Each list element will contain the patient’s name, address, phone number, and assigned doctor’s ID.
The system should allow updation of these details as well.
The system will allow a search of the doctor’s name which will return the doctor's details.
Similarly, a search of the patient’s name is also allowed which would return the patient’s details.
Similarly, a separate route should be present that takes the doctor’s ID/name as an input parameter and returns the list of patients that is handled by that doctor.
All these records should be present in the database.
Schema
Doctor

Column
Data Type
id
Unique id - char(5)
created_at
timestamp
updated_at
timestamp
name
varchar
contact_no
char(10)

Patient

Column
Data Type
id
Unique id - char(5)
created_at
timestamp
updated_at
timestamp
name
varchar
contact_no
char(10)
address
varchar(255)
doctor_id
char(5)


API Endpoints

POST /doctor/
request:
{
  "name":          ”abcd”,
  "contact_no":    “1234567890”
}

response:
{
  "id":          "10001",
  “created_at”:  1655146724,
  “updated_at”:  1655146724,
  “name”:        “abcd”,
  “contact_no” : “1234567890”
}


id should be a 5 digit uuid

GET /doctor/:id
response:
{
  "id":          "10001",
  “created_at”:  1655146724,
  “updated_at”:  1655146724,
  “name”:        “abcd”,
  “contact_no” : “1234567890”
}



PATCH /doctor/:id
request:
{
  “contact_no”:  “1234567899”
}

response:
{
  "id":          "10001",
  “created_at”:  1655146724,
  “updated_at”:  1655146724,
  “name”:  “abcd”,
  “contact_no” : “1234567899”

}
Only, “contact_no” is allowed to be updated for the doctor.


POST /patient/
request:
{
  "name":          ”xyz”,
  "contact_no":    “1234567890”,
  “address”:       “H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
  “doctor_id”:     “10001”
}

response:
{
  "id":          "20001",
  “created_at”:  1655146724,
  “updated_at”:  1655146724,
  “name”:        “xyz”,
  “contact_no” : “1234567890”,
  “address”:       “H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
  “doctor_id”:   “10001”
}
id should be a 5 digit uuid



GET /patient/:id
response:
{
  "id":          "20001",
  “created_at”:  1655146724,
  “updated_at”:  1655146724,
  “name”:        “xyz”,
  “contact_no” : “1234567890”,
  “address”:       “H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
  “doctor_id”:   “10001”
}



PATCH /patient/:id
request:
{
  “doctor_id”: “10002”
}

response:
{
  "id":          "20001",
  “created_at”:  1655146724,
  “updated_at”:  1655146724,
  “name”:        “pqrs”,
  “contact_no” : “1234567890”,
  “address”:     “H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
  “doctor_id”:   “10002”
}

Only, “contact_no”, “address” and “doctor_id” is allowed to be updated for a patient.




GET /fetchPatientByDoctorId/:docterId
response:
[
{
  "id":          "20001",
  “created_at”:  1655146724,
  “updated_at”:  1655146724,
  “name”:        “xyz”,
  “contact_no” : “1234567890”,
  “address”:     “H No.- 45/W, Street 5, Moti Bagh, West Bengal, India”
  “doctor_id”:   “10001”
},
{
  "id":          "20002",
  “created_at”:  1655146725,
  “updated_at”:  1655146725,
  “name”:        “abc”,
  “contact_no” : “1234567891”,
  “address”:     “H No.- 46/W, Street 5, Moti Bagh, West Bengal, India”
  “doctor_id”:   “10001”
}
]

This should return the list of patients being monitored by the same doctor. Here, “doctor_id” is taken in the method arguments. 


Acceptance Criteria:
Github repository with working code
Repository should have a proper Readme file with instructions for setup
Working docker container for the service                
We should be able to execute tests on the application
Extra Credit:
Implement token-based authentication
Implement mutex locking for all update/delete requests
Resources:
Tools:
https://github.com/cosmtrek/air
https://awesome-go.com/
Editors:
I personally use IntelliJ, some members use visual code as well (if you are using an IDE, use one of these)
If you are using vim, please install the go-vim plugin(https://github.com/fatih/vim-go)
Concurrency/Channels:
https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
https://github.com/devdinu/go-play#concurrency
https://divan.github.io/posts/go_concurrency_visualize/
https://www.sohamkamani.com/blog/2017/08/24/golang-channels-explained/
Interfaces:
http://www.golangbootcamp.com/book/interfaces
http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c 
https://medium.com/golangspec/interfaces-in-go-part-ii-d5057ffdb0a6
Testing:
https://segment.com/blog/5-advanced-testing-techniques-in-go/
https://medium.com/@agiratech/golang-testing-using-ginkgo-807137b0dc1e
http://onsi.github.io/ginkgo/ 
Web Development:
http://www.vividcortex.com/hubfs/eBooks/The_Ultimate_Guide_To_Building_Database-Driven_Apps_with_Go.pdf
General Guidelines and tools we use at Razorpay:
Web Development : gogin (https://github.com/gin-gonic/gin)
Logging: logrus(https://github.com/sirupsen/logrus)
ORM: Gorm(https://gorm.io/)
Testing : Built in test framework of golang
Config Management Format: toml(https://github.com/toml-lang/toml)
Config Management: BurntSushi(https://github.com/BurntSushi/toml)
Dependency /Package Management: Dep(https://github.com/golang/dep)
Dependency Injection: 
https://github.com/google/go-cloud/tree/master/wire https://blog.drewolson.org/dependency-injection-in-go 
Validation: https://github.com/go-ozzo/ozzo-validation
