<template>
  <div class="mint-page">
    <h2>铸造卡牌</h2>
    <div class="balance">当前余额: {{ balance }} GoCoin</div>

    <form @submit.prevent="handleMint" class="mint-form">
      <label>
        卡牌名字：
        <input v-model="form.name" type="text" required />
      </label>

      <label>
        投资：
        <textarea v-model="form.invest" required></textarea>
      </label>


      <button type="submit" class="button-name" role="button">铸造</button>
    </form>

    <!--overly-->
    <div v-if="showScratch || showCardModal" class="overlay">
    <!--  新增：刮刮乐 -->
      <ScratchCard
        v-if="showScratch"
        :width="450"
        :height="660"
        @complete="() => { showScratch = false; showCardModal = true }"
      >
        <!-- 刮开后显示卡牌立绘 -->
        <img 
          :src="newCard.avatar" 
          alt="new card" 
          class="scratch-card-img" 
        />
      </ScratchCard>
      <!-- 调用封装的卡牌弹窗 --> 
      <CardModal
      v-model="showCardModal"
      :card="newCard"
      :showActions="true"
      @close="closeCardModal"
    />
    </div>
  </div>
</template>

<script>
import { request } from "../api/request"
import { getProfile } from "../api/auth"
import CardModal from "../components/CardModal.vue"
import ScratchCard from "../components/ScratchToReveal.vue"   

export default {
  name: "Mint",
  components: { CardModal, ScratchCard },
  data() {
    return {
      balance: 0,
      form: { name: "", profile: "", avatar: "" },
      newCard: {},
      showCardModal: false,
      showScratch: false 
    }
  },
  async mounted() {
    await this.fetchBalance()
  },
  methods: {
    async fetchBalance() {
      try {
        const res = await getProfile()
        if (res.code === 200) {
          this.balance = res.data.gocoin
        } else {
          alert("获取余额失败")
        }
      } catch (e) {
        console.error(e)
      }
    },
    handleFileChange(e) {
      const file = e.target.files[0]
      if (!file) return
      const reader = new FileReader()
      reader.readAsDataURL(file)
      reader.onload = () => {
        this.form.avatar = reader.result
      }
    },
    async handleMint() {
      if (!this.form.name || !this.form.invest) {
        this.$root.showAlert("warning", "请填写完整信息");
        return
      }

      try {
        const investValue = parseInt(this.form.invest, 10)
        if (isNaN(investValue)) {
          this.$root.showAlert("warning", "投资必须为数字")
          return
        }
        const payload = {
          name: this.form.name,
          profile: this.form.profile,
          invest: parseInt(this.form.invest,10),
        }

        const res = await request("/auth/card/mint", {
          method: "POST",
          body: JSON.stringify(payload)
        })

        if (res.code === 200) {
          console.log("return:", res.data)
          this.newCard = res.data
          this.showScratch = true
          this.form = { name: "", profile: "", avatar  : "" }
        } else {
          this.$root.showAlert("danger", "铸造失败");
        }
      } catch (e) {
        console.error("铸造失败", e)
        alert("服务器错误")
      }
    }
  }
}
</script>

<style scoped>
.mint-page {
  max-width: 600px;
  margin: 2rem auto;
  padding: 2rem;
  background: #f9fafb;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
  text-align: center;
}

h2 {
  margin-bottom: 1rem;
  color: #1e3a8a;
}

.balance {
  margin-bottom: 1.5rem;
  font-weight: bold;
}

.mint-form label {
  display: block;
  margin-bottom: 1rem;
  text-align: left;
}

.mint-form input,
.mint-form textarea {
  width: 100%;
  padding: 0.5rem;
  margin-top: 0.3rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.button-name {
  align-items: center;
  appearance: none;
  background-color: #fcfcfd;
  border-radius: 4px;
  border-width: 0;
  box-shadow:
    rgba(45, 35, 66, 0.2) 0 2px 4px,
    rgba(45, 35, 66, 0.15) 0 7px 13px -3px,
    #d6d6e7 0 -3px 0 inset;
  box-sizing: border-box;
  color: #36395a;
  cursor: pointer;
  display: inline-flex;
  height: 48px;
  justify-content: center;
  line-height: 1;
  list-style: none;
  overflow: hidden;
  padding-left: 16px;
  padding-right: 16px;
  position: relative;
  text-align: left;
  text-decoration: none;
  transition:
    box-shadow 0.15s,
    transform 0.15s;
  user-select: none;
  -webkit-user-select: none;
  touch-action: manipulation;
  white-space: nowrap;
  will-change: box-shadow, transform;
  font-size: 18px;
}

.button-name:focus {
  box-shadow:
    #d6d6e7 0 0 0 1.5px inset,
    rgba(45, 35, 66, 0.4) 0 2px 4px,
    rgba(45, 35, 66, 0.3) 0 7px 13px -3px,
    #d6d6e7 0 -3px 0 inset;
}

.button-name:hover {
  box-shadow:
    rgba(45, 35, 66, 0.3) 0 4px 8px,
    rgba(45, 35, 66, 0.2) 0 7px 13px -3px,
    #d6d6e7 0 -3px 0 inset;
  transform: translateY(-2px);
}

.button-name:active {
  box-shadow: #d6d6e7 0 3px 7px inset;
  transform: translateY(2px);
}


.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0,0,0,0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 2000;
}
.scratch-card-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 5px; /* 和 cont_photo 保持一致 */
}

</style>
