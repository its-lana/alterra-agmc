# Deploy Golang App on Heroku

## Public URL and Repo Student Service

-  Heroku : https://students-services.herokuapp.com/status
-  Repository : https://github.com/its-lana/student-service

## How to Deploy?

-  Push Project to your Github Repository
-  Login/Signup to Heroku
-  Create new app on heroku
-  If the database is not yet hosted, you can select the `resource` tab then select the appropriate database addons
-  In the `deploy` tab, select 'Connect to Github' method then look for the appropriate repository
-  Check the option `automatic deploy`
-  Before click `Deploy Branch`, go to `settings` tab
-  In the `Config Vars` section, fill in the appropriate `.env` file for the application credentials
-  In the `Buildpacks` section select the Go language
-  You can go back to the `deploy` tab then click `Deploy Branch`
-  Done

## Screenshot

### Student API

-  Signup <br>
   ![Register](/Day-8/signup.png) <br>
-  Login <br>
   ![Login](/Day-8/login.png)<br>
-  Get All Students <br>
   ![Get All Students](/Day-8/1-get-students.png)<br>
-  Get Student by ID <br>
   ![Get Student by ID](/Day-8/1-getStudentById.png)<br>
-  Update Data Student <br>
   ![Update Data Student](/Day-8/1-updateStudent.png) <br>

### Major API

-  Get All Majors <br>
   ![Get All Majors](/Day-8/2-getMajors.png) <br>
-  Get Major By ID <br>
   ![Get Major By ID](/Day-8/2-getMajorById.png) <br>
-  Add New Major <br>
   ![Add New Major](/Day-8/2-addNewMajor.png) <br>
-  Delete Major <br>
   ![Delete Major](/Day-8/2-deleteMajor.png) <br>

### Class API

-  Get All Classes <br>
   ![Get All Classes](/Day-8/3-getClasses.png) <br>
-  Get Class By ID <br>
   ![Get Class By ID](/Day-8/3-getClassById.png) <br>
-  Add New Class <br>
   ![Add New Class](/Day-8/3-addNewClass.png) <br>
-  Update Data Class <br>
   ![Update Data Class](/Day-8/3-editClass.png) <br>
-  Delete Class <br>
   ![Delete Class](/Day-8/3-deleteClass.png) <br>
