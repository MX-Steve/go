haohui: id : 1: PHID-USER-s74srx36m7zhor3q46rf
han: id : 64: PHID-USER-ydyzkdtaughipysmt76o
ziyi: id : 11: PHID-USER-otruwkps53v2c5eiumbo
# 1. dbs:
## 1. phabricator_almanac
1. 已有表
    - almanac_binding               
    - almanac_bindingtransaction    
    - almanac_device                
    - almanac_devicename_ngrams     
    - almanac_devicetransaction     
    - almanac_interface             
    - almanac_interfacetransaction  
    - almanac_namespace             
    - almanac_namespacename_ngrams  
    - almanac_namespacetransaction  
    - almanac_network               
    - almanac_networkname_ngrams    
    - almanac_networktransaction    
    - almanac_property              
    - almanac_service               
    - almanac_servicename_ngrams    
    - almanac_servicetransaction    
    - edge                          
    - edgedata  
    - 所有表数据都为空                    
## 2. phabricator_application
1. 已有表
    - application_application          
        - 空  
    - application_applicationtransaction 
        - 
    - edge  
        - 空                             
    - edgedata                           
        - 空

## 3. phabricator_audit        
1. 已有表
    - audit_transaction           
    - audit_transaction_comment   

## 4. phabricator_auth
1. 已有表
    - auth_factorconfig              
        - 1条
    - auth_hmackey                   
        - 2条
    - auth_password     
        - 记录用户 passwordHash             
    - auth_passwordtransaction     
        - 空     
    - auth_providerconfig  
        - PhabricatorGoogleAuthProvider
        - PhabricatorLDAPAuthProvider
        - PhabricatorPasswordAuthProvider          
    - auth_providerconfigtransaction 
        - haohui/xiangfu
    - auth_sshkey    
        - 记录 sshkey                
    - auth_sshkeytransaction         
    - auth_temporarytoken   
        - 空         


## 5. phabricator_badges 
1. 已有表
    - badges_award                
    - badges_badge                
    - badges_badgename_ngrams     
    - badges_transaction          
    - badges_transaction_comment  
    - edge                        
    - edgedata                    
    - 全为空
## 6. phabricator_cache
1. 已有表
    - cache_general
    - cache_markupcache

## 7. phabricator_config
1. 已有表
    - config_entry                 
        - 7条
    - config_manualactivity        
        - 无
    - config_transaction           
        - 9

## 8. phabricator_oauth_server
1. 已有表
    - edge                                      
    - edgedata                                  
    - oauth_server_oauthclientauthorization     
    - oauth_server_oauthserveraccesstoken       
    - oauth_server_oauthserverauthorizationcode 
    - oauth_server_oauthserverclient            
    - oauth_server_transaction                  
    - 都为空

## 9. phabricator_user 





| phabricator_calendar     |
| phabricator_chatlog      |
| phabricator_conduit      |
| phabricator_conpherence  |
| phabricator_countdown    |
| phabricator_daemon       |
| phabricator_dashboard    |
| phabricator_differential |
| phabricator_diviner      |
| phabricator_doorkeeper   |
| phabricator_draft        |
| phabricator_drydock      |
| phabricator_fact         |
| phabricator_feed         |
| phabricator_file         |
| phabricator_flag         |
| phabricator_fund         |
| phabricator_harbormaster |
| phabricator_herald       |
| phabricator_legalpad     |
| phabricator_maniphest    |
| phabricator_meta_data    |
| phabricator_metamta      |
| phabricator_multimeter   |
| phabricator_nuance       |
|  |
| phabricator_owners       |
| phabricator_packages     |
| phabricator_passphrase   |
| phabricator_pastebin     |
| phabricator_phame        |
| phabricator_phlux        |
| phabricator_pholio       |
| phabricator_phortune     |
| phabricator_phragment    |
| phabricator_phrequent    |
| phabricator_phriction    |
| phabricator_phurl        |
| phabricator_policy       |
| phabricator_ponder       |
| phabricator_project      |
| phabricator_releeph      |
| phabricator_repository   |
| phabricator_search       |
| phabricator_slowvote     |
| phabricator_spaces       |
| phabricator_system       |
| phabricator_token        |
|         |
| phabricator_worker       |
| phabricator_xhpast       |
| phabricator_xhprof   