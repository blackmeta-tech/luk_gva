<template id="upload_pic">
  <el-upload
    class="avatar-uploader"
    multiple
    :http-request="handleHttpRequest"
    :show-file-list="false"
    :before-upload="beforeAvatarUpload"
    v-loading.fullscreen.lock="loading"
    element-loading-text="拼命上传中"
    element-loading-background="rgba(0,0,0,0.5)"
    accept=".jpeg,.png,.jpg"
  >
    <img v-if="imageUrl" :src="imageUrl" class="avatar">
    <el-icon v-if="!imageUrl" class="avatar-uploader-icon"><plus /></el-icon>
  </el-upload>
</template>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
<script>
import axios from 'axios'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage } from 'element-plus'

const userStore = useUserStore()
export default {
  props: ['pic', 'picKey'],
  data() {
    return {
      loading: false,
      imageUrl: this.pic,
      pKey: this.picKey,
      lastFile: {
        key: '',
        url: '',
      }
    }
  },
  watch: {
    pic(imageUrl) {
      this.imageUrl = imageUrl
    },
    picKey(pKey) {
      this.pKey = pKey
    }
  },
  methods: {
    handleHttpRequest(file) {
      this.loading = true
      var fd = new FormData()
      fd.append('file', file.file)
      axios.post('/api/fileUpload/upload', fd, {
        headers: {
          'Content-Type': 'multipart/form-data',
          'x-token': userStore.token,
          'x-user-id': userStore.userInfo.id,
        }
      }).then(response => {
        if (response.data.code === 0) {
          this.$emit('updatePic', response.data.data.file.url)
          this.$emit('updatePicKey', response.data.data.file.key)
          this.imageUrl = response.data.data.file.url
          this.lastFile = response.data.data.file
          this.loading = false
          ElMessage({ type: 'success', message: '上传成功' })
        } else {
          this.loading = false
          ElMessage({
            showClose: true,
            message: response.data.msg,
            type: 'error'
          })
        }
      })
    },

    beforeAvatarUpload(file) {
      const isLt2M = file.size / 1024 / 1024 < 2
      if (!isLt2M) {
        // this.$message.error('上传头像图片大小不能超过 2MB!')
      }
      return isLt2M
    },

    // 删除文件
    deleteFile(data) {
      if (data.key) {
        axios.post('/api/fileUpload/delete', data, {
          headers: {
            'Content-Type': 'application/json',
            'x-token': userStore.token,
            'x-user-id': userStore.userInfo.id,
          }
        })
      }
    }
  }
}
</script>
