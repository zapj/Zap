<template>
  
    <table class="el-table__header">
        <thead>
            <tr>
                <td></td>
                <td></td>
                <td></td>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>CPU 名称</td>
                <td>{{ state.serverInfo.cpu_name }}</td>
                <td></td>
            </tr>
            <tr>
                <td>CPU 频率</td>
                <td>{{ state.serverInfo.cpu_ghz }}GHz</td>
                <td></td>
            </tr>
            <tr>
                <td>CPU 核心数</td>
                <td>{{ state.serverInfo.cpu_cores }}核心</td>
                <td></td>
            </tr>
        </tbody>
    </table>
  </template>
  
  <script setup>
import { onMounted,ref,reactive } from 'vue';
import apiRequest from '../../httpclient/client';
  const state = reactive({
        serverInfo:{}
    })
  

  onMounted(()=>{
    getServerInfo()
  })

  const getServerInfo = ()=>{
    apiRequest({
        url:"/v1/server/info"
    }).then((resp)=>{
        state.serverInfo = resp.data
        console.log(state.serverInfo);
        
    })
  }
  </script>
  