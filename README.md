## SECTION 1 : PROJECT TITLE
## 4M1L - Credit Card Recommendation System

<img src="Miscellaneous/home.png"
     style="float: left; margin-right: 0px;" />

---
## SECTION 2 : EXECUTIVE SUMMARY / PAPER ABSTRACT
A vibrant city located in the heart of Asia, Singapore offers global investors unparalleled access to global markets. Strategically located to serve Asia Pacific, one of the world’s fastest-growing regions, Singapore’s well-established business infrastructure, global connectivity and trade linkages enable investors to access the approximately 4 billion strong Asian market within a radius of 7 hours’ flight time.
In Singapore, one of the first things people do after getting their first job is to apply for a credit card. It makes perfect sense that paying with a credit card allows them to enjoy many benefits that using cash wouldn’t provide. With the advancement of technology, the people nowadays could even embed the credit card with their phone with the various mobile payment introduced such as Apple Pay, Samsung Pay, Google Pay and etc.

In the past, carrying credit cards were considered one of the elusive “5Cs” (Cash, Car,
Credit card, Condominium and Country club membership) that people in Singapore
worked hard to attain. Today, owning a credit card is not only the norm, but the
financially-savvy option.

However, picking a credit card can be a very frustrating process and most of the people
have went through a hard time in finding the best credit card. By trying to optimize the
reward from credit card, they might have fallen into the pitfall of applying excessive
credit card, resulting in burdened themselves with unnecessary late payment charge or
principal annual fee. Overall, no single credit card is better than all others in all
categories, but by asking the right questions, each person can find the card that best fit his spending habit and credit situation.

Our team, comprising of 5 members from different nationalities, looking forward to
achieve financial freedom in one day. As the first step towards the financial freedom, we
have developed a credit card recommendation system, with the aim to help all the first
time credit card applicants in optimizing the credit card selection and better their
financial management.


---
## SECTION 3 : CREDITS / PROJECT CONTRIBUTION

| Official Full Name  | Student ID (MTech Applicable)  | Work Items (Who Did What) | Email (Optional) |
| :------------ |:---------------:| :-----| :-----|
| Chen Liwei | A0101217B | Video Editing Data Collection Knowledge Modelling| e0384319@u.nus.edu |
| Lee Boon Kien | A0195175W | Video Presentation Data Collection Knowledge Modelling| e0384806@u.nus.edu |
| Ng Cheong Hong| A0195290Y| Knowledge Modelling Rules Engine Programming | e0384921@u.nus.edu |
| Raymond Djajalaksana| A0195381X | Report Writing Front End Programming | e0385012@u.nus.edu |
| Seah Jun Ru| A0097451Y | Data Collection Report Writing Knowledge Modelling| e0258166@u.nus.edu |

---
## SECTION 4 : VIDEO OF SYSTEM MODELLING & USE CASE DEMO

[![Credit Card Recommendation System](http://img.youtube.com/vi/kF0tPmweUeU/0.jpg)](http://www.youtube.com/watch?v=kF0tPmweUeU)

---
## SECTION 5 : USER GUIDE

Requirements:
* nodejs and npm should be installed. Otherwise please download and install from the following website: https://www.npmjs.com/get-npm
* Python 2.7 ~ Python 3.6.5 should be installed. Otherwise, please download and install from the following website: https://www.python.org/downloads/ 
[*It is recommended to use the python version mentioned above since durable_rules might not be working properly for other python version]

Installation:
- [Anaconda](https://repo.anaconda.com/archive/Anaconda3-2018.12-Windows-x86_64.exe "Anaconda") / [Python 3.6](https://www.python.org/downloads/release/python-365/ "Python 3.6") or older
- [Node.js ](https://nodejs.org/en/ "Node.js ")
- Microsoft Visual C++ 14.0: [Build Tools for Visual Studio 2017](https://visualstudio.microsoft.com/thank-you-downloading-visual-studio/?sku=BuildTools&rel=15 "Build Tools for Visual Studio 2017")  

``` bash
# 1. install all front end dependencies
cd web/
npm install

# 2. install all backend dependencies
pip install requests flask flask_cors durable_rules

# 3. (Windows only) run start.bat to start all application 
./start.bat
# 3. (Non windows) you need to run redis server manually, and then run the rules engine by running cc_system.py inside rules-engine folder 
python cc_system.py


# 4. Try access localhost:8080/home 


# Alternatively you just need to run rules engine by running python script and redis
start_server.bat
# And then access the frontend from AWS host. This static host will connect to your localhost rules engine backend
http://machine-reasoning.s3-website-ap-southeast-1.amazonaws.com/home

```

User Guide
`4M1L_User_Guide_CCRS.pdf` : <https://github.com/raycap/IRS-MR-2019-01-19-IS1PT-GRP-4M1L-CCRS/blob/master/UserGuide/4M1L_User_Guide_CCRS.pdf>

---
## SECTION 6 : PROJECT REPORT / PAPER
`4M1L_CreditCardRecommendationReport.pdf` : <https://github.com/raycap/IRS-MR-2019-01-19-IS1PT-GRP-4M1L-CCRS/blob/master/ProjectReport/4M1L_CreditCardRecommendationReport.pdf>

---
## SECTION 7 : MISCELLANEOUS

### [Credit card selection survey-2.csv](https://github.com/raycap/IRS-MR-2019-01-19-IS1PT-GRP-4M1L-CCRS/blob/master/Miscellaneous/Credit%20card%20selection%20survey-2.csv)
* Results of survey
* Insights derived, which were subsequently used in our system
### 

### [Credit Card Database.csv](https://github.com/raycap/IRS-MR-2019-01-19-IS1PT-GRP-4M1L-CCRS/blob/master/Miscellaneous/Credit%20Card%20Database.xlsx)
* Selection criterias for credit cards, which were collated from various banks.
* Served as the basis to calculate for the eligible cashflow amount
### 

---



---

**This [Machine Reasoning (MR)](https://www.iss.nus.edu.sg/executive-education/course/detail/machine-reasoning "Machine Reasoning") course is part of the Analytics and Intelligent Systems and Graduate Certificate in [Intelligent Reasoning Systems (IRS)](https://www.iss.nus.edu.sg/stackable-certificate-programmes/intelligent-systems "Intelligent Reasoning Systems") series offered by [NUS-ISS](https://www.iss.nus.edu.sg "Institute of Systems Science, National University of Singapore").**

**Lecturer: [GU Zhan (Sam)](https://www.iss.nus.edu.sg/about-us/staff/detail/201/GU%20Zhan "GU Zhan (Sam)")**

[![alt text](https://www.iss.nus.edu.sg/images/default-source/About-Us/7.6.1-teaching-staff/sam-website.tmb-.png "Let's check Sam' profile page")](https://www.iss.nus.edu.sg/about-us/staff/detail/201/GU%20Zhan)

**zhan.gu@nus.edu.sg**
