
#Proof of Work implementation
#Author:    Bastian Buck
#Date:      08.10.2019


from hashlib import sha256, sha512, md5                                                             
from random import randint                                                                          

def pow(text):                                                                                      #simple with sha256 hashes and 10 binary nulls                                                                
    found = False                                                                                   
    nonce = randint(1000,100000)                                                                    

    while (found==False):                                                                           
        hasht = sha256((text+str(nonce)).encode()).hexdigest()                                      #make sha256 hash with text + random nonce into hex-format
        #print(hasht)        #optional output of all generated hashes                            
                                                                                                    
        if(hasht.startswith("002") or hasht.startswith("003")):                                     #proof about 10 binary nulls 
                                                                                                    
            found = True                                                                             
            print("Hash found! "+hasht)                                                          
            print("Text and Nonce: "+text+" + "+str(nonce)+"\n")                                    
        nonce = nonce+1                                                                             
#pow("bstnbuck")                                                                                  



def pown():                                                                                         #implementation with selection between md5, sha256 and sha512                                                         
    found = False                                                                                   
    nonce = randint(1000,100000)    
    count = 0
    text = input("Please type in name: ")
    nulls = input("Type in leading nulls (in hex! like: 002 = 10 binary nulls): ")
    try:
        mode = int(input("Choose mode: SHA256 = 1, SHA512 = 2, MD5 = 3 (hashing method not safe): "))
    except:
        print("ERROR! Only numbers allowed")
        return
    while (found==False):  
        if(mode==1):
            hasht = sha256((text+str(nonce)).encode()).hexdigest()
            #print(hasht)        #optional output of all generated hashes                                                                                                                           
            if(hasht.startswith(nulls)):                                                                                                                          
                found = True                                                                           
                print("Hash found! "+str(hasht))                                                          
                print("Text and Nonce: "+text+" + "+str(nonce)+" count: "+str(count)+" \n")
            nonce +=1 
            count +=1

        elif(mode==2):
            hasht = sha512((text+str(nonce)).encode()).hexdigest()
            #print(hasht)        #optional output of all generated hashes                                                                                                                         
            if(hasht.startswith(nulls)):                                                                                                                          
                found = True                                                                           
                print("Hash found! "+str(hasht))                                                          
                print("Text and Nonce: "+text+" + "+str(nonce)+" count: "+str(count)+"\n")
            nonce+=1  
            count+=1

        elif(mode==3):
            hasht = md5((text+str(nonce)).encode()).hexdigest()
            #print(hasht)        #optional output of all generated hashes                                                                                                                        
            if(hasht.startswith(nulls)):                                                                                                                          
                found = True                                                                           
                print("Hash found! "+str(hasht))                                                          
                print("Text and Nonce: "+text+" + "+str(nonce)+" count: "+str(count)+"\n")
            nonce+=1  
            count+=1

        else:
            print("ERROR! Wrong Mode!")
            found = True

def start():
    starti = False
    
    while(starti==False):
        try:                                                                                    #catch errors by typing in 
            startinput = int(input("Start Proof of Work? 1= Yes, 9= No: "))
            if(startinput == 1):
                pown()
            elif(startinput == 9):
                print("Stopped!")
                starti = True
            else:
                print("ERROR! Wrong entry")
        except:
            print("ERROR!")
       
start()

def proof(hasht, text):                                                                         #function to proof pow
    if(hasht == sha256(text.encode()).hexdigest()):                                                 
        print("Hash correct! "+hasht)                                                              
    else:                                                                                           
        print("Hash incorrect! "+hasht)                                                             


#proof("00000001c4cef3c4ac540a27c5e9a242926cb6cee82b6d79db06ca4d1ef6ff73","buck12852971")
#proof("0003cb9eaaf9b95729753079e5b74f7de9ecd8d4fd2b0a670a0041dcbeae2537","something247599")