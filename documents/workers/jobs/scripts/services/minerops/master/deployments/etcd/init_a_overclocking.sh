#!/usr/bin/env bash
# exm: 1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875;1100/875,2000/875
# 全局参数
export ETCDCTL_API=3
ETCDCTL="etcdctl --endpoints=127.0.0.1:2379 "
DC_UUID="3f154da2-4532-4435-abe6-758fc4f7fc85"
global_key="/shepherd/${DC_UUID}/worker/occonfig"
IFname="ens33"
IFvalue=$(ip a  | grep -C 1 $IFname | awk '/ether/{print $2}')
mac_key="/shepherd/${DC_UUID}/worker/$IFvalue/occonfig"
LogFile="/tmp/overclocking.log"
> $LogFile

#declare -A card0 card1 card2 card3 card4 card5 card6 card7
card0=()
card1=()
card2=()
card3=()
card4=()
card5=()
card6=()
card7=()

ifDataOk(){
  value=$1
  f=";" # 7
  d="," # 8
  x="/" # 16
  f_num=`grep -o "$f" <<< "$value" | wc -l`
  d_num=`grep -o "$d" <<< "$value" | wc -l`
  x_num=`grep -o "$x" <<< "$value" | wc -l`
  if [[ $f_num != 7 ]] || [[ $d_num != 8 ]] || [[ $x_num != 16 ]];then
    echo "数据存储结构不对，请检查"
    echo "数据存储结构不对，请检查" >> $LogFile
    exit 2
  fi
}

checkEtcdctlExists(){
  $ETCDCTL version &>/dev/null
  if [ $? != 0 ];then
    echo "Etcdctl 在客户端上没有安装"
    echo "Etcdctl 在客户端上没有安装" >> $LogFile
    exit 1
  fi
}

# 解析字符串，获取对应值
getKv(){
  config=$1
  cards=(${config//;/ })
  i=0
  for card in ${cards[@]}
  do
     sms=(${card//,/ })
     s=${sms[0]}
     m=${sms[1]}
     if [ $i == 0 ];then
       card0=(${s/\//\ } "m 2 "${m/\//\ })
     elif [ $i == 1 ]; then
       card1=(${s/\//\ } "m 2 "${m/\//\ })
     elif [ $i == 2 ]; then
       card2=(${s/\//\ } "m 2 "${m/\//\ })
     elif [ $i == 3 ]; then
       card3=(${s/\//\ } "m 2 "${m/\//\ })
     elif [ $i == 4 ]; then
       card4=(${s/\//\ } "m 2 "${m/\//\ })
     elif [ $i == 5 ]; then
       card5=(${s/\//\ } "m 2 "${m/\//\ })
     elif [ $i == 6 ]; then
       card6=(${s/\//\ } "m 2 "${m/\//\ })
     elif [ $i == 7 ]; then
       card7=(${s/\//\ } "m 2 "${m/\//\ })
     fi
     i=$[$i+1]
  done
}

setOverclocking(){
  for i in 0 1 2 3 4 5 6 7
  do
    echo "set card$i"
    echo "set card$i"  >> $LogFile
    if [ $i == 0 ];then
      s=${card0[0]}" "${card0[1]}
      m=${card0[2]}" "${card0[3]}" "${card0[4]}" "${card0[5]}
    elif [ $i == 1 ];then
      s=${card1[0]}" "${card1[1]}
      m=${card1[2]}" "${card1[3]}" "${card1[4]}" "${card1[5]}
    elif [ $i == 2 ];then
      s=${card2[0]}" "${card2[1]}
      m=${card2[2]}" "${card2[3]}" "${card2[4]}" "${card2[5]}
    elif [ $i == 3 ];then
      s=${card3[0]}" "${card3[1]}
      m=${card3[2]}" "${card3[3]}" "${card3[4]}" "${card3[5]}
    elif [ $i == 4 ];then
      s=${card4[0]}" "${card4[1]}
      m=${card4[2]}" "${card4[3]}" "${card4[4]}" "${card4[5]}
    elif [ $i == 5 ];then
      s=${card5[0]}" "${card5[1]}
      m=${card5[2]}" "${card5[3]}" "${card5[4]}" "${card5[5]}
    elif [ $i == 6 ];then
      s=${card6[0]}" "${card6[1]}
      m=${card6[2]}" "${card6[3]}" "${card6[4]}" "${card6[5]}
    elif [ $i == 7 ];then
      s=${card7[0]}" "${card7[1]}
      m=${card7[2]}" "${card7[3]}" "${card7[4]}" "${card7[5]}
    fi
    echo "echo manual > /sys/class/drm/card$i/device/power_dpm_force_performance_level"
    echo "echo s 0 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo s 1 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo s 2 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo s 3 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo s 4 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo s 5 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo s 6 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo s 7 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo $m > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo c > /sys/class/drm/card$i/device/pp_od_clk_voltage"
    echo "echo 3 > /sys/class/drm/card$i/device/pp_dpm_sclk"
    echo "echo 2 > /sys/class/drm/card$i/device/pp_dpm_mclk"
    echo "echo manual > /sys/class/drm/card$i/device/power_dpm_force_performance_level" >> $LogFile
    echo "echo s 0 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo s 1 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo s 2 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo s 3 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo s 4 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo s 5 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo s 6 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo s 7 $s > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo $m > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo c > /sys/class/drm/card$i/device/pp_od_clk_voltage" >> $LogFile
    echo "echo 3 > /sys/class/drm/card$i/device/pp_dpm_sclk" >> $LogFile
    echo "echo 2 > /sys/class/drm/card$i/device/pp_dpm_mclk" >> $LogFile
  done
}

main(){
  checkEtcdctlExists
  # 1. 获取局部 cards
  mac_config=$($ETCDCTL get $mac_key)
  mac_config=${mac_config: 80}
  if [ -n "$mac_config" ]; then
    ifDataOk $mac_config
    echo "使用局部调优配置"
    echo "使用局部调优配置" >> $LogFile
    getKv $mac_config
    setOverclocking
  else
    # 2. 如果局部配置为空，则使用全局配置
    global_config=$($ETCDCTL get $global_key)
    global_config=${global_config: 63}
    if [ -n "$global_config" ]; then
      ifDataOk $global_config
      echo "使用全局调优配置"
      echo "使用全局调优配置" >> $LogFile
      getKv $global_config
      setOverclocking
    else
      echo "当前 DC 还没有全局默认超频参数设置"
      echo "当前 DC 还没有全局默认超频参数设置" >> $LogFile
    fi
  fi
}
main
