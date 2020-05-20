const record = {
  state: {
    taskId: '',
    recordIndex: 0
  },
  mutations: {
    SET_TASK_ID: (state, taskId) => {
      state.taskId = taskId
    },
    SET_RECORD_INDEX: (state, recordIndex) => {
      state.recordIndex = recordIndex
    }
  },
  actions: {}
}

export default record
