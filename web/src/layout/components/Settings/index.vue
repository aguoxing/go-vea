<template>
  <el-drawer v-model="showSettings" :withHeader="false" direction="rtl" size="300px">
    <el-divider content-position="center">主题</el-divider>
    <div class="theme-switch">
      <el-switch v-model="isDark" inline-prompt active-icon="ep:moon" inactive-icon="ep:sunny" @change="toggleDark" />
    </div>
    <el-divider content-position="center">导航栏模式</el-divider>
    <div class="setting-drawer-block-checbox">
      <div class="setting-drawer-block-checbox-item" @click="handleTheme('theme-dark')">
        <img src="@/assets/images/dark.svg" alt="dark" />
        <div v-if="sideTheme === 'theme-dark'" class="setting-drawer-block-checbox-selectIcon" style="display: block">
          <i aria-label="图标: check" class="anticon anticon-check">
            <svg viewBox="64 64 896 896" data-icon="check" width="1em" height="1em" :fill="theme" aria-hidden="true" focusable="false" class>
              <path
                d="M912 190h-69.9c-9.8 0-19.1 4.5-25.1 12.2L404.7 724.5 207 474a32 32 0 0 0-25.1-12.2H112c-6.7 0-10.4 7.7-6.3 12.9l273.9 347c12.8 16.2 37.4 16.2 50.3 0l488.4-618.9c4.1-5.1.4-12.8-6.3-12.8z"
              />
            </svg>
          </i>
        </div>
      </div>
      <div class="setting-drawer-block-checbox-item" @click="handleTheme('theme-light')">
        <img src="@/assets/images/light.svg" alt="light" />
        <div v-if="sideTheme === 'theme-light'" class="setting-drawer-block-checbox-selectIcon" style="display: block">
          <i aria-label="图标: check" class="anticon anticon-check">
            <svg viewBox="64 64 896 896" data-icon="check" width="1em" height="1em" :fill="theme" aria-hidden="true" focusable="false" class>
              <path
                d="M912 190h-69.9c-9.8 0-19.1 4.5-25.1 12.2L404.7 724.5 207 474a32 32 0 0 0-25.1-12.2H112c-6.7 0-10.4 7.7-6.3 12.9l273.9 347c12.8 16.2 37.4 16.2 50.3 0l488.4-618.9c4.1-5.1.4-12.8-6.3-12.8z"
              />
            </svg>
          </i>
        </div>
      </div>
    </div>
    <el-divider content-position="center">主题色</el-divider>
    <div>
      <el-color-picker v-model="theme" :predefine="predefineColors" @change="themeChange" />
    </div>
    <el-divider content-position="center">系统布局配置</el-divider>
    <div class="drawer-item">
      <span>开启 TopNav</span>
      <span>
        <el-switch v-model="topNav" class="drawer-switch" />
      </span>
    </div>

    <div class="drawer-item">
      <span>开启 Tags-Views</span>
      <span>
        <el-switch v-model="tagsView" class="drawer-switch" />
      </span>
    </div>

    <div class="drawer-item">
      <span>固定 Header</span>
      <span>
        <el-switch v-model="fixedHeader" class="drawer-switch" />
      </span>
    </div>

    <div class="drawer-item">
      <span>显示 Logo</span>
      <span>
        <el-switch v-model="sidebarLogo" class="drawer-switch" />
      </span>
    </div>

    <div class="drawer-item">
      <span>动态标题</span>
      <span>
        <el-switch v-model="dynamicTitle" class="drawer-switch" />
      </span>
    </div>

    <el-divider />

    <el-button type="primary" plain @click="saveSetting"> <ep:document /> 保存配置 </el-button>
    <el-button plain @click="resetSetting"> <ep:refresh /> 重置配置 </el-button>
  </el-drawer>
</template>

<script lang="ts" setup>
import variables from '@/assets/styles/variables.module.scss'
import originElementPlus from 'element-plus/theme-chalk/index.css'
import { useDynamicTitle } from '@/utils/dynamicTitle'
import { handleThemeStyle } from '@/utils/theme'
import { localStorage } from '@/utils/storage'

import useStore from '@/store'
import { useDark, useToggle } from '@vueuse/core'
const { app, setting, permission } = useStore()

const isDark = useDark()
const toggleDark = useToggle(isDark)

const { proxy } = getCurrentInstance()
const showSettings = ref(false)
const theme = ref(setting.theme)
const sideTheme = ref(setting.sideTheme)
const storeSettings = computed(() => setting)
const predefineColors = ref(['#409EFF', '#ff4500', '#ff8c00', '#ffd700', '#90ee90', '#00ced1', '#1e90ff', '#c71585'])

/** 是否需要topnav */
const topNav = computed({
  get: () => setting.topNav,
  set: val => {
    setting.changeSetting({ key: 'topNav', value: val })
    if (!val) {
      app.toggleSideBarHide(false)
      permission.setSidebarRouters(permission.defaultRoutes)
    }
  }
})
/** 是否需要tagview */
const tagsView = computed({
  get: () => storeSettings.value.tagsView,
  set: val => {
    setting.changeSetting({ key: 'tagsView', value: val })
  }
})
/**是否需要固定头部 */
const fixedHeader = computed({
  get: () => storeSettings.value.fixedHeader,
  set: val => {
    setting.changeSetting({ key: 'fixedHeader', value: val })
  }
})
/**是否需要侧边栏的logo */
const sidebarLogo = computed({
  get: () => storeSettings.value.sidebarLogo,
  set: val => {
    setting.changeSetting({ key: 'sidebarLogo', value: val })
  }
})
/**是否需要侧边栏的动态网页的title */
const dynamicTitle = computed({
  get: () => storeSettings.value.dynamicTitle,
  set: val => {
    setting.changeSetting({ key: 'dynamicTitle', value: val })
    // 动态设置网页标题
    useDynamicTitle()
  }
})

function themeChange(val) {
  setting.changeSetting({ key: 'theme', value: val })
  theme.value = val
  handleThemeStyle(val)
}
function handleTheme(val) {
  setting.changeSetting({ key: 'sideTheme', value: val })
  sideTheme.value = val
}
function saveSetting() {
  proxy.$modal.loading('正在保存到本地，请稍候...')
  let layoutSetting = {
    topNav: storeSettings.value.topNav,
    tagsView: storeSettings.value.tagsView,
    fixedHeader: storeSettings.value.fixedHeader,
    sidebarLogo: storeSettings.value.sidebarLogo,
    dynamicTitle: storeSettings.value.dynamicTitle,
    sideTheme: storeSettings.value.sideTheme,
    theme: storeSettings.value.theme
  }
  localStorage.set('layout-setting', JSON.stringify(layoutSetting))
  setTimeout(proxy.$modal.closeLoading(), 1000)
}
function resetSetting() {
  proxy.$modal.loading('正在清除设置缓存并刷新，请稍候...')
  localStorage.remove('layout-setting')
  setTimeout('window.location.reload()', 1000)
}
function openSetting() {
  showSettings.value = true
}

defineExpose({
  openSetting
})
</script>

<style lang="scss" scoped>
.setting-drawer-title {
  margin-bottom: 12px;
  color: rgba(0, 0, 0, 0.85);
  line-height: 22px;
  font-weight: bold;
  .drawer-title {
    font-size: 14px;
  }
}
.setting-drawer-block-checbox {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  margin-top: 10px;
  margin-bottom: 20px;

  .setting-drawer-block-checbox-item {
    position: relative;
    margin-right: 16px;
    border-radius: 2px;
    cursor: pointer;

    img {
      width: 48px;
      height: 48px;
    }

    .custom-img {
      width: 48px;
      height: 38px;
      border-radius: 5px;
      box-shadow: 1px 1px 2px #898484;
    }

    .setting-drawer-block-checbox-selectIcon {
      position: absolute;
      top: 0;
      right: 0;
      width: 100%;
      height: 100%;
      padding-top: 15px;
      padding-left: 24px;
      color: #1890ff;
      font-weight: 700;
      font-size: 14px;
    }
  }
}

.drawer-item {
  color: rgba(0, 0, 0, 0.65);
  padding: 12px 0;
  font-size: 14px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.theme-switch {
  text-align: center;
}
</style>
