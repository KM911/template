<template>
  <nut-cell
    :showIcon="true"
    title="选择单个日期"
    :desc="state.date ? `${state.date} ${state.dateWeek}` : '请选择'"
    @click="openSwitch('isVisible')"
  >
  </nut-cell>
  <nut-calendar
    v-model:visible="state.isVisible"
    :default-value="state.date"
    @close="closeSwitch('isVisible')"
    @choose="setChooseValue"
    :start-date="`2022-01-11`"
    :end-date="`2022-11-30`"
  >
  </nut-calendar>
</template>
<script setup>
import { reactive } from 'vue';
const state = reactive({
  isVisible: false,
  date: '',
  dateWeek: ''
});
const openSwitch = (param) => {
  state[`${param}`] = true;
};
const closeSwitch = (param) => {
  state[`${param}`] = false;
};
const setChooseValue = (param) => {
  state.date = param[3];
  state.dateWeek = param[4];
};
</script>
