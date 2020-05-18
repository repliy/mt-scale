const task = {
  state: {
    taskId: ''
  },
  getters: {
    getTaskId: state => {
      return state.taskId
    }
  },
  mutations: {
    SET_TASK_ID: (state, taskId) => {
      state.taskId = taskId
    }
  },
  actions: {}
}

export default task
