<script setup lang="ts">
import { ref } from 'vue'
import Tips from './components/tips.vue'

const searchValue = ref('')
const isEmpty = ref(true)
const loading = ref(false)
const result = ref<any | null>(null)
const error = ref<string | null>(null)

function validateInput(val: string) {
  const pattern = /^(?:\d{5,6}|T01\d{4}|JC[48]\d{4}|JY01\d{4}|SJ01\d{4})$/i
  return pattern.test(val)
}

function onInput(e: Event) {
  let val = (e.target as HTMLInputElement).value

  
  val = val.replace(/[^a-zA-Z0-9]/g, '')

  
  val = val.toUpperCase()

  
  val = val.slice(0, 8)

  
  error.value = null
}


async function handleSearch() {
  const val = searchValue.value.trim()

  if (!val) {
    error.value = '请输入车厢号'
    result.value = null
    return
  }

  if (!validateInput(val)) {
    error.value = '输入格式不正确，请检查车厢号'
    result.value = null
    return
  }

  error.value = null
  result.value = null
  loading.value = true

  try {
    const res = await fetch('https://<dns(not ip)>:port/api/calculate', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ carriage_number: val }),
    })

    if (!res.ok) throw new Error('查询失败')

    const data = await res.json()

    if (data.StateCode === '2000') {
      if (Array.isArray(data.Data) && data.Data.length > 0) {
        result.value = data.Data
      } else if (Array.isArray(data.Data) && data.Data.length === 0) {
        result.value = null
        error.value = data.Msg || '未找到相关车厢'
      } else {
        result.value = null
        error.value = data.Msg || '返回数据格式不正确'
      }
    } else if (data.StateCode === '5000') {
      result.value = null
      if (data.Data && data.Data.Msg) {
        error.value = data.Data.Msg
      } else {
        error.value = data.Msg || '服务器处理错误'
      }
    } else {
      result.value = null
      error.value = data.Msg || '未知错误'
    }
  } catch (e: any) {
    error.value = e.message || '查询出错'
  } finally {
    loading.value = false
  }
}
</script>



<template>
  <div class="container">
    <h1>上海地铁车厢号数据库</h1>

    <div class="search-bar">
      <input
        v-model="searchValue"
        type="text"
        maxlength="8"
        placeholder="请输入车厢号"
        @input="onInput"
        @keyup.enter="handleSearch"
        autocomplete="off"
      />
      <button :disabled="loading" @click="handleSearch"><span>Go!</span></button>
    </div>
    <transition name="tips-fade">
      <Tips v-if="!searchValue" />
    </transition>

    <!-- Loading / 结果区 -->
    <div class="result" v-if="loading || result">
      <template v-if="loading">
        <span>查询中...</span>
      </template>

      <!-- 多条结果（数组） -->
      <template v-else-if="Array.isArray(result)">
        <div v-for="(train, idx) in result" :key="idx" class="train-item">
          <!-- 当 IsInputCarriageCorrect 为 false 时，显示正确的输入提示 -->
          <div v-if="train.IsInputCarriageCorrect === false" class="error" style="margin-bottom:10px;">
            输入车厢号有误，正确应该为
            <strong>
              {{
                (() => {
                  const arr = Array.isArray(train.Carriage_number) ? train.Carriage_number : null;
                  const idxNum = Number(train.Carriage_index);
                  if (arr && Number.isFinite(idxNum) && arr[idxNum] !== undefined) {
                    return arr[idxNum];
                  }
                  // 回退：如果没有索引或数组，展示整个 Carriage_number 或提示
                  return arr ? arr.join('、') : (train.Carriage_num || '未知');
                })()
              }}
            </strong>
          </div>

          <div>车号：<strong>{{ train.TrainId }}</strong></div>
          <div>车型代码：<strong>{{ train.Train_type }}</strong></div>
          <div>车厢数量：<strong>{{ Array.isArray(train.Carriage_number) ? train.Carriage_number.length : '未知' }}</strong></div>
          <div>关于该列车：{{ train.TrainDetail || '暂无信息' }}</div>

          <div style="margin-top:8px;">
            <small>车厢号列表：
              <span v-if="Array.isArray(train.Carriage_number)">
                {{ train.Carriage_number.join('、') }}
              </span>
              <span v-else>无</span>
            </small>
          </div>

          <hr v-if="idx < result.length - 1" />
        </div>
      </template>

      <!-- 单条对象结果（备用） -->
      <template v-else-if="typeof result === 'object' && result">
        <div class="train-item">
          <div v-if="result.IsInputCarriageCorrect === false" class="error" style="margin-bottom:8px;">
            输入车厢号有误，正确应该为
            <strong>
              {{
                 (() => {
                  const arr = Array.isArray(result.Carriage_number) ? result.Carriage_number : null;
                  const idxNum = Number(result.Carriage_index);
                  if (arr && Number.isFinite(idxNum) && arr[idxNum] !== undefined) return arr[idxNum];
                  return arr ? arr.join('、') : (result.Carriage_num || '未知');
                })()
              }}
            </strong>
          </div>

          <div>车号：<strong>{{ result.TrainId }}</strong></div>
          <div>车型代码：<strong>{{ result.Train_type }}</strong></div>
          <div>关于该列车：{{ result.TrainDetail || '暂无信息' }}</div>
          <div style="margin-top:8px;">
            <small>车厢号列表：
              <span v-if="Array.isArray(result.Carriage_number)">{{ result.Carriage_number.join('、') }}</span>
              <span v-else>无</span>
            </small>
          </div>
        </div>
      </template>
    </div>

    <!-- 错误提示 -->
    <div class="error" v-if="error && !loading">{{ error }}</div>
    <footer class="footer">
      <p>
        发现 Bug 或有改进建议？请前往
        <a
          href="https://github.com/Martyr404/shmetroDB/issues"
          target="_blank"
          rel="noopener noreferrer"
        >
          GitHub Issues
        </a>
      </p>
    </footer>
  </div>
</template>


<style scoped>
:root {
  --theme-color: #0078d7;
  --theme-light: #e8f3ff;
  --error-color: #e74c3c;
  --text-color: #222;
  --border-color: #d0d0d0;
  --bg-white: #fff;
  --radius: 12px;
  --shadow: 0 4px 24px rgba(0,0,0,0.08);
}

.container {
  max-width: 520px;
  margin: 60px auto;
  padding: 36px 28px;
  border-radius: 1rem;
  box-shadow: var(--shadow);
  font-family: 'Segoe UI', 'PingFang SC', Arial, sans-serif;
  color: var(--text-color);
  text-align: center;
  transition: all 0.3s ease;
  box-shadow: 0 8px 32px rgba(0, 120, 215, 0.15);
}


h1 {
  font-size: 1.5rem;
  margin-bottom: 28px;
  color: var(--theme-color);
  letter-spacing: 1px;
  font-weight: 600;
}

.search-bar {
  display: flex;
  gap: 10px;
  justify-content: center;
  margin-bottom: 24px;
}

input[type="text"] {
  flex: 1;
  padding: 10px 14px;
  border: 1.5px solid var(--border-color);
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.25s, box-shadow 0.25s;
}

input[type="text"]:focus {
  border-color: var(--theme-color);
  box-shadow: 0 0 6px rgba(0,120,215,0.25);
  outline: none;
}

button {
  background-color: transparent;
  border: 2px solid #409eff;
  color: #409eff;
  font-size: 1rem;
  font-weight: 600;
  padding: 10px 24px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.25s ease;
  outline: none;
  letter-spacing: 0.5px;
}


button:hover {
  background-color: #8fc5fb;
  border: 2px solid#8fc5fb;
  color: #fff;
  box-shadow: 0 4px 10px rgba(64, 158, 255, 0.25);
  transform: translateY(-2px);
}


button:active {
  transform: translateY(0);
  box-shadow: 0 2px 6px rgba(64, 158, 255, 0.2);
}

button:disabled {
  border-color: #cfd9e0;
  color: #cfd9e0;
  background-color: transparent;
  cursor: not-allowed;
  box-shadow: none;
}





.result {
  margin-top: 24px;
  text-align: left;
}

.train-item {
  background: var(--theme-light);
  border-left: 4px solid var(--theme-color);
  border-radius: 10px;
  padding: 14px 16px;
  margin-bottom: 14px;
  box-shadow: 0 2px 6px rgba(0,0,0,0.04);
  transition: all 0.25s ease;
}

.train-item:hover {
  background: #f5faff;
  box-shadow: 0 3px 12px rgba(0,120,215,0.12);
  transform: translateY(-2px);
}

.train-item div {
  margin: 4px 0;
  font-size: 0.96rem;
}

.train-item small {
  color: #555;
  font-size: 0.85rem;
}

.result hr {
  border: none;
  border-bottom: 1px dashed #ccc;
  margin: 10px 0;
}

.error {
  margin-top: 22px;
  color: var(--error-color);
  background: #fde8e8;
  border-left: 4px solid var(--error-color);
  padding: 10px 12px;
  border-radius: 6px;
  font-size: 0.96rem;
  text-align: left;
}
.footer {
  margin-top: 48px;
  padding-top: 18px;
  border-top: 1px solid rgba(0, 0, 0, 0.08);
  font-size: 0.92rem;
  color: #666;
  text-align: center;
  line-height: 1.8;
  user-select: none;
}

.footer a {
  color: #409eff;
  font-weight: 600;
  text-decoration: none;
  position: relative;
  transition: color 0.25s ease, text-shadow 0.25s ease;
}

.footer a::after {
  content: "";
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 100%;
  height: 1.5px;
  background: linear-gradient(90deg, #4e9fff, #6b66ff);
  transform: scaleX(0);
  transform-origin: right;
  transition: transform 0.3s ease;
}


.footer a:hover::after {
  transform: scaleX(1);
  transform-origin: left;
}


/* 响应式：手机端 */
@media (max-width: 600px) {
  .container {
    padding: 24px 16px;
    margin: 30px auto;
  }
  .search-bar {
    flex-direction: column;
    gap: 10px;
    align-items: stretch;
  }
  input[type="text"] {
    width: 100%;
    box-sizing: border-box;
  }
  button {
    width: 100%;
    box-sizing: border-box;
  }
  .train-item {
    font-size: 0.9rem;
  }
}

/* tip进入和离开都使用过渡 */
.tips-fade-enter-active,
.tips-fade-leave-active {
  transition: opacity 0.35s ease;
}

/* tip初始状态：透明 */
.tips-fade-enter-from,
.tips-fade-leave-to {
  opacity: 0;
}

/* tip完整显示 */
.tips-fade-enter-to,
.tips-fade-leave-from {
  opacity: 1;
}
</style>

