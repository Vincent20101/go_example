<template>
    <div class="login">
        <a-card class="login-card" :headStyle="{padding: '0px'}" :bodyStyle="{padding: '20px'}">
            <template #title>
                <div class="login-card-header">
                    <img style="height:25px;margin:0 8px 5px 0" :src="kubeLogo"/>
                    <span>KubeA</span>
                </div>
            </template>
            <a-form ref="formRef" :model="loginData" :label-col="{span: 5}">
                <a-form-item
                    label="用户名"
                    name="username"
                    :rules="[{required: true, message: '请输入账号'}]">
                    <a-input v-model:value="loginData.username"></a-input>
                </a-form-item>
                <a-form-item
                    label="密码"
                    name="password"
                    :rules="[{required: true, message: '请输入密码'}]">
                    <a-input-password v-model:value="loginData.password"></a-input-password>
                </a-form-item>
                <a-form-item style="text-align:center;margin-bottom:10px;">
                    <c-button style="width:100%" type="primary" size="normal" icon="user-outlined" @click="onCheck()">登录</c-button>
                </a-form-item>
            </a-form>
        </a-card>
    </div>
</template>

<script>
import kubeLogo from '@/assets/k8s-metrics.png';
import { reactive, ref } from 'vue';
import common from "@/config";
import { useRouter } from 'vue-router';
import moment from 'moment';
import jwt from 'jsonwebtoken';
import httpClient from '@/request';
import { message } from 'ant-design-vue';
export default({
    setup() {
        const formRef = ref()
        const loginData = reactive({
            username: '',
            password: ''
        })
        const loginUrl = common.loginAuth
        const router = useRouter()
        //登录方法
        //这种方式是同步执行的方法
        const onCheck = async () => {
            try {
                //具体同步执行的内容
                await formRef.value.validateFields()
                handleLogin()
            } catch (errorInfo) {
                console.log('Failed:', errorInfo)
            }
        }
        function handleLogin() {
            httpClient.post(loginUrl, loginData)
            .then(res => {
                localStorage.setItem('username', loginData.username)
                localStorage.setItem('loginDate', moment().format('YYYY-MM-DD HH:mm:ss'))
                //生成token
                let token = jwt.sign(loginData, 'adoodevops', { expiresIn: '10h' });
                localStorage.setItem('token', token)
                router.push('/')
                message.success(res.msg)
            })
            .catch(res => {
                message.error(res.msg)
            })
        }
        return {
            kubeLogo,
            formRef,
            loginData,
            onCheck
        }
    },
})
</script>

<style scoped>
/* 添加scoped，表示页面之间的css互不干扰 */
    /* 调整背景及布局 */
    .login {
        height: 100vh;
        background-image: url(../../assets/login3.webp);
        background-size: 100%;
        display: grid;
        place-items: center;
    }
    .login-card {
        width: 350px;
    }
    .login-card-header {
        font-size: 18px;
        font-weight: bold;
        text-align: center;
    }
</style>