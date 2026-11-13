Lab-4 : Mini CTF
•	I pasted https://www.google.com/search?q=
%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38
%62%33%65%32%33%39%65%62%63%32&rlz=1C5CHF into google and found this string

•	This string does not look like bcrypt as what the challenge said since bcrypt would start with $2b$, so I assume it would be MD5 since it is shorter than what bcrypt would be
 
<img width="1280" height="660" alt="image" src="https://github.com/user-attachments/assets/28f4bfac-e12e-4d55-a1ac-90f73a284f3a" />

•	Then I paste that string to CrackStation and find the password is meowmeow and the hash type is MD5
•	After that I use regex to check if it was correct or not as the lab suggest
 
<img width="1250" height="662" alt="image" src="https://github.com/user-attachments/assets/f048436c-3b87-47b0-af76-684cace6a448" />

<img width="1896" height="790" alt="Screenshot 2026-11-13 204248" src="https://github.com/user-attachments/assets/eebae3cf-c06d-4319-86f2-f3bae2c47968" />

	 
