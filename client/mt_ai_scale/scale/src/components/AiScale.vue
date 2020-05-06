<template>
  <div>
    <div id="config">
      <div class="select d-inline-block">
        <label for="videoSource">Video source: </label>
        <select id="videoSource"></select>
      </div>
      <div class="select d-inline-block">
        <label for="modelSource">Model: </label>
        <select id="base_model">
          <option value="mobilenet_v1">Bayshore Scale_30k</option>
          <option value="mobilenet_v2">Bayshore Scale_100k</option>
        </select>
      </div>
    </div>
    <div id="camera1">
      <video
        class="size"
        autoPlay
        playsInline
        muted
        ref="videoRef"
        width="640"
        height="480"
      />
      <canvas class="size" ref="canvasRef" width="640" height="480" />
    </div>
  </div>
</template>

<script>
import * as cocoSsd from '@tensorflow-models/coco-ssd'
import '@tensorflow/tfjs'
import { filterValidPredictions } from '../utils/digits_reader.js'
import { returnWeight } from '../utils/digits_reader.js'

export default {
  data() {
    return {
      video: null,
      canvas: null,
      videoSources: [],
      videoInputDevices: [],
    }
  },
  mounted() {
    this.video = this.$refs.videoRef
    this.canvas = this.$refs.canvasRef
    // this.getStream().then(this.getDevices).then(this.gotDevices)
    //  var canvas = this.$refs.canvasRef
    this.getDevices().then(this.gotDevices)
  },
  methods: {
    getDevices() {
      return navigator.mediaDevices.enumerateDevices()
    },
    gotDevices(deviceInfos) {
      for (const device of deviceInfos) {
        if (device.kind === 'videoinput') {
          this.videoInputDevices.push(device)
        }
      }
      console.log(this.videoInputDevices)
      const modelPromise = cocoSsd.load({
        base: 'mobilenet_v1',
      })
      const webCamPromise = this.getStream()
      Promise.all([modelPromise, webCamPromise]).then((result) => {
        this.detectFrame(result[0])
      })
    },
    detectFrame(model) {
      model.detect(this.video).then((predictions) => {
        const ctx = this.canvas.getContext('2d')
        ctx.clearRect(0, 0, ctx.canvas.width, ctx.canvas.height)
        console.log(ctx.canvas.width)
        // Font options.
        const font = '16px sans-serif'
        ctx.font = font
        ctx.textBaseline = 'top'

        predictions
          .filter((prediction) => prediction.class !== 'Bad Number')
          .filter((prediction) => prediction.score > 0.1)
          .forEach((prediction) => {
            const x = prediction.bbox[0]
            const y = prediction.bbox[1]
            const width = prediction.bbox[2]
            const height = prediction.bbox[3]
            // Draw the bounding box.
            ctx.strokeStyle = '#00FFFF'
            ctx.lineWidth = 4
            ctx.strokeRect(x, y, width, height)
            // Draw the label background.
            ctx.fillStyle = '#00FFFF'
            const label_text =
              prediction.class + ' ' + prediction.score.toFixed(2)
            const textWidth = ctx.measureText(label_text).width
            const textHeight = parseInt(font, 10) // base 10
            ctx.fillRect(x, y, textWidth + 4, textHeight + 4)
            // Draw the text last to ensure it's on top.
            ctx.fillStyle = '#000000'
            ctx.fillText(label_text, x, y)
            console.log(label_text)
          })
        const result = filterValidPredictions(predictions)
        // call set reading function
        // returnWeight(result);
        console.log('result: ', result)
        requestAnimationFrame(() => {
          this.detectFrame(model)
        })
      })
    },
    getStream() {
      if (window.stream) {
        window.stream.getTracks().forEach((track) => {
          track.stop()
        })
      }
      const myDeviceId = this.videoInputDevices[0].deviceId
      const constraints = {
        video: { deviceId: myDeviceId ? { exact: myDeviceId } : undefined },
      }
      return navigator.mediaDevices
        .getUserMedia(constraints)
        .then(this.gotStream)
        .catch(this.handleError)
    },
    gotStream(stream) {
      window.stream = stream
      this.video.srcObject = stream
      return new Promise((resolve, reject) => {
        this.video.onloadedmetadata = () => {
          resolve()
        }
      })
    },
    handleError(error) {
      console.log('Error: ', error)
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.size {
  position: absolute;
  top: 10;
  left: 0;
}
</style>
