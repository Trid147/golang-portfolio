#!/bin/bash

#Logs imitating bash script

NORMAL_IPS=("192.168.1.50" "192.168.1.51" "172.16.0.5" "10.0.0.15" "192.168.1.102")
ATTACK_IPS=("10.0.0.99" "198.51.100.42" "203.0.113.11")

LEGAL_PAGES=("/index.html" "/about.html" "/contact.html" "/pricing.html" "/services.html")
LEGAL_ASSETS=("/css/styles.css" "/js/bundle.js" "/images/hero.jpg" "/favicon.ico")

while true; do
  TIMESTAMP=$(date '+%d/%b/%Y:%H:%M:%S %z')
  CHANCE=$((1 + RANDOM % 12))
  
  if [ $CHANCE -le 4 ]; then
    #Legal standard page request
    IP=${NORMAL_IPS[$((RANDOM % ${#NORMAL_IPS[@]}))]}
    PAGE=${LEGAL_PAGES[$((RANDOM % ${#LEGAL_PAGES[@]}))]}
    SIZE=$((2000 + RANDOM % 6000))
    echo "$IP - - [$TIMESTAMP] \"GET $PAGE HTTP/1.1\" 200 $SIZE" >> mock_access.log

  elif [ $CHANCE -le 6 ]; then
    #Legal standard asset request
    IP=${NORMAL_IPS[$((RANDOM % ${#NORMAL_IPS[@]}))]}
    ASSET=${LEGAL_ASSETS[$((RANDOM % ${#LEGAL_ASSETS[@]}))]}
    STATUS=$(STATUSES=(200 304); echo ${STATUSES[$((RANDOM % 2))]})
    SIZE=$([ "$STATUS" = "200" ] && echo $((400 + RANDOM % 2500)) || echo "0")
    echo "$IP - - [$TIMESTAMP] \"GET $ASSET HTTP/1.1\" $STATUS $SIZE" >> mock_access.log

  elif [ $CHANCE -eq 7 ]; then
    #Legal product view
    IP=${NORMAL_IPS[$((RANDOM % ${#NORMAL_IPS[@]}))]}
    echo "$IP - - [$TIMESTAMP] \"GET /products?id=5 HTTP/1.1\" 200 4324" >> mock_access.log

  elif [ $CHANCE -eq 8 ]; then
    #SQL Injection
    IP=${ATTACK_IPS[$((RANDOM % ${#ATTACK_IPS[@]}))]}
    echo "$IP - - [$TIMESTAMP] \"GET /products?id=5' UNION SELECT username, password FROM users-- HTTP/1.1\" 403 8940" >> mock_access.log

  elif [ $CHANCE -eq 9 ]; then
    #Path Traversal
    IP=${ATTACK_IPS[$((RANDOM % ${#ATTACK_IPS[@]}))]}
    echo "$IP - - [$TIMESTAMP] \"GET /download?file=../../../../etc/passwd HTTP/1.1\" 404 230" >> mock_access.log

  elif [ $CHANCE -eq 10 ]; then
    #Cross-Site Scripting
    IP=${ATTACK_IPS[$((RANDOM % ${#ATTACK_IPS[@]}))]}
    echo "$IP - - [$TIMESTAMP] \"GET /search?q=<script>alert(1)</script> HTTP/1.1\" 403 180" >> mock_access.log

  elif [ $CHANCE -eq 11 ]; then
    #Remote Command Execution
    IP=${ATTACK_IPS[$((RANDOM % ${#ATTACK_IPS[@]}))]}
    echo "$IP - - [$TIMESTAMP] \"POST /cgi-bin/test.cgi HTTP/1.1\" 500 530" >> mock_access.log

  else
    #Additional legal traffic
    IP=${NORMAL_IPS[$((RANDOM % ${#NORMAL_IPS[@]}))]}
    LEGAL_EXTRA=$((1 + RANDOM % 2))
    if [ $LEGAL_EXTRA -eq 1 ]; then
      echo "$IP - - [$TIMESTAMP] \"GET /search?q=security+logging HTTP/1.1\" 200 3120" >> mock_access.log
    else
      echo "$IP - - [$TIMESTAMP] \"POST /contact-submit HTTP/1.1\" 302 0" >> mock_access.log
    fi
  fi
  
  sleep $((1 + RANDOM % 6))
done
