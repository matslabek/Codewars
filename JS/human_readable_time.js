//https://www.codewars.com/kata/52685f7382004e774f0001f7/train/javascript
function humanReadable(seconds) {
    let hours = Math.floor(seconds / 3600);
    let minutes = Math.floor(seconds / 60) - (hours * 60);
    let secs = seconds - (hours * 3600) - (minutes * 60);

    let timeArray = [hours, minutes, secs];
    let timeString = "";
    timeArray.forEach(function(value) {
        if (value < 10){
            timeString += "0";
        }
        timeString += String(value) + ":";
    });
    return timeString.slice(0,-1);
}