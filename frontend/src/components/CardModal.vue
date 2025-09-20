<template>
  <!-- 只有 visible 为 true 时才显示 -->
   <div>
   <Alert ref="alertRef" />
  <div class="cont_principal" v-if="modelValue" @click="onBackgroundClick">
    <div class="cont_central">
      <div :class="['cont_modal', { cont_modal_active: isActive }]" @click.stop>
        <!-- 卡牌图片及信息 -->
        <div class="cont_photo">
          <div class="cont_img_back">
            <img :src="card.avatar" alt="card image" />
          </div>

        </div>

        <!-- 卡牌详细内容 -->
        <div class="cont_text_ingredients">
          <div class="cont_over_hidden">
            <div class="cont_tabs">
              <!-- From Uiverse.io by Voxybuns --> 
              <button>
                <span class="button_top" @click="toggleSale(card)">{{ card.onsale ? '下架' : '上架' }}</span>
              </button>
              <!-- From Uiverse.io by Voxybuns --> 
              <button>
                <span class="button_top" @click="destroyCard(card)">{{ '销毁' }}</span>
              </button>

            </div>

            <div class="cont_text_det_preparation">
              <div class="cont_title_preparation">
              </div>
              <div class="cont_info_preparation">
                <p>卡牌名称：{{ card.name }}</p>
                <div class="card-attributes">
                  <span :style="{ color: getRarityColor(card.rarity) }">
                    稀有度: {{ card.rarity }}
                  </span>
                  <span>费用: {{ card.cost }}</span>
                  <span>攻击力: {{ card.attack }}</span>
                  <span>血量: {{ card.blood }}</span>
                  <span>{{ card.hashid }}</span>
                </div>

              </div>

              <div class="cont_title_preparation">
                <p>Description</p>
              </div>
              <div class="cont_info_preparation">
                <p>{{ card.profile }}</p>
              </div>
            </div>
          </div>

        </div>


      </div>
    </div>
  </div>
  </div>
</template>

<script setup>
import { request } from "../api/request";
import { ref, onMounted, watch,inject } from "vue";
import Alert from '@/components/Alert.vue';
const alertRef = ref(null); // 用来调用 Alert 的方法
const emit = defineEmits(['update:modelValue','close']);
const showAlert = inject("showAlert"); // ✅ 这样才能拿到 provide 的方法
// 父组件通过props传入卡牌数据和显示状态
const props = defineProps({
  card: { type: Object, default: () => ({}) },
  modelValue: {
    type: Boolean,
    default: false
  }
});

const isActive = ref(false);

let autoOpenTimeout = null; // <- 定义定时器变量

watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) {
      if (autoOpenTimeout) clearTimeout(autoOpenTimeout);
      autoOpenTimeout = setTimeout(() => {
        isActive.value = true;
      }, 0); // 延迟 500ms 自动展开
    } else {
      isActive.value = false;
      if (autoOpenTimeout) clearTimeout(autoOpenTimeout);
    }
  },
  { immediate: true }
);


// 切换展开/收缩
function toggleActive() {
  isActive.value = !isActive.value;
}
function onBackgroundClick() {
  isActive.value = false;
  emit("update:modelValue", false);
  emit('close');
  console.log(111);
}

function getRarityColor(rarity) {
  switch(rarity) {
    case 'common': return '#999';      // 灰色
    case 'rare': return '#1E90FF';    // 蓝色
    case 'epic': return '#8A2BE2';    // 紫色
    case 'legendary': return '#FF8C00';    // 橙色
    default: return '#565656';       // 默认颜色
  }
}



// 调用父组件或全局方法（最小改动）
async function destroyCard(card) {
  if (!confirm("确定要销毁这张卡牌吗？")) return;
  try {
    const res = await request("/auth/card/destroy", {
      method: "POST",
      body: JSON.stringify({ hashid: card.hashid })
    });
    if (res.code === 200) {
      showAlert("success","销毁成功");
      window.location.reload();
    } else alert(res.msg);
  } catch (e) {
    console.error(e);
    alert("销毁失败");
  }
}

async function toggleSale(card) {
  try {
    if (!card.onsale) {
      let inputPrice = prompt("请输入上架价格(GoCoin)", "10");
      if (inputPrice === null) return;
      inputPrice = parseInt(inputPrice);
      if (isNaN(inputPrice) || inputPrice <= 0) { alert("请输入有效数字"); return; }
      const res = await request("/auth/card/sell", {
        method: "POST",
        body: JSON.stringify({ hashid: card.hashid, cost: inputPrice })
      });
      if (res.code === 200) { card.onsale = true;showAlert("success","上架成功"); window.location.reload(); }
      else alert("上架失败");
    } else {
      const res = await request("/auth/card/cancel", {
        method: "POST",
        body: JSON.stringify({ orderid: card.transid })
      });
      if (res.code === 200) { card.onsale = false;showAlert("success","下架成功"); window.location.reload(); }
      else alert("下架失败");
    }
  } catch (e) { console.error(e); alert("操作失败"); }
}

</script>

<style scoped>
/* -------------- Basic Reset Css ---------- */
* {
  margin: 0px auto;
  padding: 0px;
  text-align: center;
  font-family: 'Open Sans', sans-serif;
}

/* 去掉背景，只保留布局 */
.cont_principal {
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  z-index:999;
  background: rgba(0,0,0,0.3); /* 可选，仅用于点击区域可见调试 */

}

.cont_central {  
  position: absolute;
  top:50%;
  left:50%;
  transform: translate(-50%, -50%);
}

.cont_modal {
  position: relative;
  width: 450px;
  height: 600px;
  transition: all 0.5s;
  transition-delay: 0.7s;
}

.cont_photo {
  position: relative;
  width: 450px;
  height: 660px;
  overflow: hidden;
  background-color: #eee;
  border-radius:5px;
  top: -20px;
  float: left;
  z-index: 2;
  transition: all 0.5s;
  box-shadow: 1px 1px 20px -5px rgba(0,0,0,0.5);  
}

.cont_img_back {
  position: absolute;
  width: 100%;
  height: 100%;
  overflow: hidden;
  border-radius:5px ;
}

.cont_img_back > img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  opacity: 0.7;
  transition: all 1s;
}


/* 以下 CSS 保留原样动画和布局 */
.cont_text_ingredients {
  position: absolute;
  width: 0px;
  top: 0px;
  left: 435px;
  margin-left: 10px;
  height: 600px;
  float: left; 
  border-radius:5px ;
  z-index: 3;
  box-shadow: 1px 1px  20px  -5px  rgba(0,0,0,0.2);
  background: linear-gradient(135deg, rgba(251,249,249,1) 28%,rgba(232,234,237,1) 100%);
  transition: all 0.5s; 
  transition-delay: 0.7s;
}

.cont_mins{
  position: relative;
  float: left;
  width: 100%;
}



.cont_detalles {
  position:absolute;
  bottom:-185px;
  height:200px;
  border-radius:5px;
  background: linear-gradient(to bottom, rgba(0,0,0,0) 0%, rgba(0,0,0,0.65) 68%, rgba(0,0,0,0.65) 100%);
  width:100%;
  transition:all 0.5s;
  transition-delay:0.7s;
}

.cont_detalles > h3 { margin-top:50px; color:#fff; font-size:24px; }
.cont_detalles > p { color:#fff; width:80%; text-align:left; font-size:14px; }

/* Tabs, preparation steps and button animations 保留原样 */
.cont_tabs { position:relative; float:left; width:610px; height:70px; display: flex; justify-content: space-between; align-items: center; padding: 0 10px; border-bottom:3px solid #EDEDEC;}
.cont_tabs button {
  flex: 1;                /* 平分可用空间 */
  margin: 0 10px;          /* 按钮间距 */
}
.cont_tabs button .button_top {
  width: 100%;            /* 让按钮宽度填满 flex */
  text-align: center;     /* 文字居中 */
}

.cont_title_preparation { position:relative; float:left; margin:10px 0px; width:410px; }
.cont_title_preparation > p { font-weight:700; font-size:14px; margin-left:40px; text-align:left; color:#36354E; }

.cont_info_preparation { position:relative; float:left; }
.cont_info_preparation > p { margin:5px 0px; margin-left:50px; border-left:2px solid #E3E3E3; font-size:18px; font-weight: 600; padding:20px 0px 20px 20px; text-align:left; padding-right:15px; color:#565656; }

.cont_over_hidden { position:relative; float:left; width:100%; height:400px; overflow:hidden; }
.cont_text_det_preparation { position:relative; width:410px; }

.cont_modal_active > .cont_text_ingredients > .cont_btn_open_dets > a { transform:rotate(0deg);}
.cont_modal_active > .cont_text_ingredients{ width:615px; left:427px; z-index:1; box-shadow:15px 20px 70px -5px rgba(0,0,0,0.2);}
.cont_modal_active { width:1050px; }
.cont_modal_active > .cont_photo { box-shadow:25px 10px 70px -5px rgba(0,0,0,0.3);}
.cont_modal_active > .cont_photo > .cont_mins > .sub_mins { opacity:1; }
.cont_modal_active > .cont_photo > .cont_servings { opacity:1; }
.cont_modal_active > .cont_photo > .cont_detalles { bottom:0px; }

.card-attributes {
  display: flex;
  flex-wrap: wrap;       /* 自动换行 */
  justify-content: flex-start;
  gap: 20px 40px;        /* 行间距20px，列间距40px */
  margin-left: 50px;     /* 保持和原先 Step 内容对齐 */
}

.card-attributes span {
  font-size: 12px;
  color: #565656;
  font-weight: 700; /* 加粗文字 */

}

/* From Uiverse.io by Voxybuns */ 
button {
  /* Variables */
  --button_radius: 0.75em;
  --button_color: #e8e8e8;
  --button_outline_color: #000000;
  font-size: 17px;
  font-weight: bold;
  border: none;
  cursor: pointer;
  border-radius: var(--button_radius);
  background: var(--button_outline_color);
}

.button_top {
  display: block;
  box-sizing: border-box;
  border: 2px solid var(--button_outline_color);
  border-radius: var(--button_radius);
  padding: 0.75em 1.5em;
  background: var(--button_color);
  color: var(--button_outline_color);
  transform: translateY(-0.2em);
  transition: transform 0.1s ease;
}

button:hover .button_top {
  /* Pull the button upwards when hovered */
  transform: translateY(-0.33em);
}

button:active .button_top {
  /* Push the button downwards when pressed */
  transform: translateY(0);
}


</style>
