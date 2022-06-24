<template>
    <form action="">
        <h4>
            标题
        </h4>
        <div>
            <input type="text" v-model="title" class="title" placeholder="请输入标题" />
        </div>
        <h4>
            留言内容
        </h4>
        <div>
            <textarea placeholder="请输入留言内容" v-model="content"></textarea>
        </div>
        <div class="submit-btn-box">
            <button class="submit-btn" type="button" @click="submitContent">提交</button>
        </div>
    </form>
    <ul class="comment-list">
        <li v-for="(item, index) in comments" :key="index">
            <h5>{{item.title}}</h5>
            <div class="content">
                {{item.content}}
            </div>
            <div class="time-box">
                <time> {{item.create_time}} </time>
            </div>
        </li>
    </ul>
</template>

<script>
export default {
    name: "CommentIndex",
    data() {
        return {
            url: "http://127.0.0.1:8080",
            content: "",
            title: "",
            comments: []
        }
    },
    methods: {
        submitContent(){
            var formData = new FormData();
            formData.set("title", this.title)
            formData.set("content", this.content)

            fetch(
                this.url + "/comment/submit", 
                {
                    method: "POST",
                    body: formData
                }
            ).then(response => {
                if(response.status == 200 ){
                    return response.json();
                }
                return {}
            }).then(data=>{
                if(data.status == 200) {
                    alert("发表成功");
                    location.reload();
                }
            }).catch(err=>{
                console.log(err);
            })
        },
    },
    created() {
        fetch(this.url+"/comment/list").then(response=>{
            if(response.status == 200) {
                return response.json()
            }
            return {}
        }).then(data=>{
            if(data.status == 200){
                this.comments = data.data
            }
        }).catch(err => {
            console.log(err);
        })
    },
}
</script>

<style scoped>  
    form {
        display: block;
        width: 400px;
        margin: 50px auto;
    }
    .title {
        width: 380px;
        padding: 10px;
        background-color: #e7e7e7;
        border: none;
        outline: none;
        border-radius: 5px;
    }
    textarea {
        width: 380px;
        height: 100px;
        background-color: #e7e7e7;
        border: none;
        outline: none;
        border-radius: 5px;
        padding: 10px;
        resize: none;
        font-family:sans-serif
    }
    .submit-btn {
        border: none;
        background-color: rgb(104, 104, 189);
        color: white;
        border-radius: 5px;
        padding: 5px 20px;

    }
    .submit-btn-box {
        text-align: right;
    }

    .comment-list {
        width: 400px;
        margin: 60px auto;
        padding: 0;
    }

    .comment-list li {
        list-style: none;
        padding-top: 20px;
    }
    .time-box {
        color: gray;
        text-align: right;
        font-size: 12px;
        padding-top: 10px;
    }
    .content {
        font-size: 13px;
        text-indent: 28px;
    }
</style>