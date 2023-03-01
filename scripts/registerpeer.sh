#!/usr/bin/env bash
#参数1组织uscc, 参数2 peer名,名字和密码使用同一参数, 参数3 peer域名, 参数4 ca端口
export FABRIC_CA_CLIENT_HOME=/root/txhyjuicefs/organizations/fabric-ca/$1/

infoln "Register peer"
set -x
fabric-ca-client register -d --caname ca-$1 --id.name $2 --id.secret $2 --id.type peer --tls.certfiles /root/txhyjuicefs/organizations/fabric-ca/$1/ca-cert.pem
{ set +x; } 2>/dev/null

mkdir -p /root/txhyjuicefs/organizations/$1/peers
mkdir -p /root/txhyjuicefs/organizations/$1/peers/$3

infoln "Generate the peer msp"
set -x
fabric-ca-client enroll -d -u https://$2:$2@localhost:$4 --caname ca-$1 -M /root/txhyjuicefs/organizations/$1/peers/$3/msp --csr.hosts $3 --tls.certfiles /root/txhyjuicefs/organizations/fabric-ca/$1/ca-cert.pem
{ set +x; } 2>/dev/null

cp /root/txhyjuicefs/organizations/$1/msp/config.yaml /root/txhyjuicefs/organizations/$1/peers/$3/msp/config.yaml

infoln "Generate the peer tls"
set -x
fabric-ca-client enroll -d -u https://$2:$2@localhost:$4 --caname ca-$1 -M /root/txhyjuicefs/organizations/$1/peers/$3/tls --enrollment.profile tls --csr.hosts $3 --csr.hosts localhost --tls.certfiles /root/txhyjuicefs/organizations/fabric-ca/$1/ca-cert.pem
{ set +x; } 2>/dev/null

cp /root/txhyjuicefs/organizations/$1/peers/$3/tls/tlscacerts/* /root/txhyjuicefs/organizations/$1/peers/$3/tls/ca.crt
cp /root/txhyjuicefs/organizations/$1/peers/$3/tls/signcerts/* /root/txhyjuicefs/organizations/$1/peers/$3/tls/server.crt
cp /root/txhyjuicefs/organizations/$1/peers/$3/tls/keystore/* /root/txhyjuicefs/organizations/$1/peers/$3/tls/server.key
