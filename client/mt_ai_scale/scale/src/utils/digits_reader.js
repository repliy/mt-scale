// digits reading model handling
export function check_overlap(box1, box2) {
  // box[x,y,w,h]
  return (((box1[0] >= box2[0] && box1[0] <= box2[0] + box2[2]) ||
    (box1[0] + box1[2] >= box2[0] && box1[0] + box1[2] <= box2[0] + box2[2])) &&
  ((box1[1] >= box2[1] && box1[1] <= box2[1] + box2[3]) ||
    (box1[1] + box1[3] >= box2[1] && box1[1] + box1[3] <= box2[1] + box2[3])))
}

export function filterValidPredictions(predictions) {
  console.log('+++++++' + JSON.stringify(predictions) + '\n')

  const num_of_precision = 2
  const average = arr => arr.reduce((sume, el) => sume + el, 0) / arr.length
  var result = ''
  // // find digital screen poi first
  // const screen_boxes = predictions
  //   .filter(prediction =>
  //     (prediction.score > 0.6 && prediction.class === "Digital Screen"))
  //   // find the largest one
  //   .sort(function(a,b) {
  //     return a.bbox[2] * a.bbox[3] - b.bbox[2] * b.bbox[3];
  //   })

  // const screen_box = screen_boxes[0]
  const good_numbers = predictions
    .filter(prediction =>
    //      (check_overlap(prediction,screen_box) &&
      (prediction.score > 0.8) // && prediction.class !== "Digital Screen")
    )
  // sort from right to left on x axis
    .sort(function(a, b) { return a.bbox[0] - b.bbox[0] })

  const avg_number_width = average(good_numbers)
  // get number starting from the lowest digit
  // get second digits after dot
  var digit_count = good_numbers.length - 1

  for (var i = good_numbers.length - 1; i >= 0; i--) {
    result = good_numbers[i].class + result
    if (i === good_numbers.length - 2) {
      result = '.' + result
    }
  }
  return result
}

export function returnWeight(result) {
  console.log('--------' + result + '----------')
  var scale = document.getElementById('initweight').value
  if (result != scale) {
    // Show Lobster Size
    var reverse = document.getElementById('reverse').value
    if (reverse == '1') {
      lobstersize(result, scale)
    }
    document.getElementById('reading_1').value = result
    document.getElementById('reading_1').onchange()
    document.getElementById('initweight').value = result
  }
}

// Calculate the weight of the lobster according to the change of the weighing value
// before and after, and display the corresponding information.
export function lobstersize(result, scale) {
  var weight = parseFloat(scale) - parseFloat(result)
  document.getElementById('lob_weight').value = weight
  if (parseFloat(weight) > 0.8 & parseFloat(weight) <= 1.25) {
    document.getElementById('sizediv').style.background = '#FFFFFF'
    document.getElementById('sizediv').style.color = '#ED008C'
    document.getElementById('sizediv').innerHTML = 'CHIX'
  }
  if (parseFloat(weight) > 1.25 & parseFloat(weight) <= 1.5) {
    document.getElementById('sizediv').style.background = '#FFFFFF'
    document.getElementById('sizediv').style.color = '#00A885'
    document.getElementById('sizediv').innerHTML = 'Quarter'
  }
  if (parseFloat(weight) > 1.75 & parseFloat(weight) <= 2) {
    document.getElementById('sizediv').style.background = '#FFFFFF'
    document.getElementById('sizediv').style.color = '#02A5E8'
    document.getElementById('sizediv').innerHTML = 'Select'
  }
  if (parseFloat(weight) > 2 & parseFloat(weight) <= 2.5) {
    document.getElementById('sizediv').style.background = '#FFFFFF'
    document.getElementById('sizediv').style.color = '#5F2C91'
    document.getElementById('sizediv').innerHTML = 'S D'
  }
  if (parseFloat(weight) > 2.5 & parseFloat(weight) <= 3) {
    document.getElementById('sizediv').style.background = '#5F2C91'
    document.getElementById('sizediv').style.color = '#FFFFFF'
    document.getElementById('sizediv').innerHTML = 'L D'
  }
  if (parseFloat(weight) > 3 & parseFloat(weight) <= 4) {
    document.getElementById('sizediv').style.background = '#FFFFFF'
    document.getElementById('sizediv').style.color = '#FFE01B'
    document.getElementById('sizediv').innerHTML = '3 - 4'
  }
  if (parseFloat(weight) > 4 & parseFloat(weight) <= 6) {
    document.getElementById('sizediv').style.background = '#FFFFFF'
    document.getElementById('sizediv').style.color = '#01A99C'
    document.getElementById('sizediv').innerHTML = '4 - 6'
  }
  if (parseFloat(weight) > 6) {
    document.getElementById('sizediv').style.background = '#FFFFFF'
    document.getElementById('sizediv').style.color = '#231F20'
    document.getElementById('sizediv').innerHTML = '6 +'
  }
  if (parseFloat(weight) < 0) {
    document.getElementById('sizediv').style.background = '#FFFFFF'
    document.getElementById('sizediv').style.color = 'red'
    document.getElementById('sizediv').innerHTML = 'ERROR'
  }
}
