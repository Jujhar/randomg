package randomg

import (
 "time"
 "strconv"
)

func CreateInteger(numDigits int) string {
  // fmt.Printf("You want %d digits\n", numDigits);
  var milli = (time.Now().UnixNano() / int64(time.Millisecond));
  sub := strconv.FormatInt(milli, 10)

  // get last 4 string for random purposes
  last4  := sub[len(sub)-5:]
  last4 = last4[:len(last4)-1]

  // return the correct number of digits
  if (numDigits <= 4) {
      x := 4 - numDigits
      return(last4[:len(last4)-x])
  } else {

    // get last 3 string for random purposes
    last3  := sub[len(sub)-4:];
    last3 = last3[:len(last3)-1];

    num3, _ := strconv.ParseInt(last3, 10, 64)

    // Generate long enough numbers
    for (numDigits > len(strconv.FormatInt(num3, 10))) {
        num3 = num3 * 3;
	  }

    result := strconv.FormatInt(num3, 10) + last4;

    // return the correct number of numDigits
    result = result[0:numDigits];

    return(result);
  }
}
