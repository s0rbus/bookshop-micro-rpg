var shopOpened = false
var shopDaycount = 0
var currentDay = 0
var verbose = false

function setVerbose(v) {
    verbose = v
}

function getName() {
    return "simple, coffee shop expansion"
}

function getRequiredThrows() {
    return 3
}

function run(day, throws) {
    var req = getRequiredThrows()
    if(throws.length != req) {
        throw `you must provide ${req} dice throws.`
    }
    var runStatus = ""
    if(shopOpened && shopDaycount > 0) {
        runStatus = "running"
    }
    var action = []
    //1 in 3 chance that shop opens if not already opened
    if(throws[0] < 3) {
        if(!shopOpened) {
            shopOpened = true
            shopDaycount = 1
            currentDay = day
            runStatus = "opening"
            let action1 = {score:-7,category:"MONEY",description:"Bought stock for shop"}
            action.push(JSON.stringify(action1))
            let action2 = {score:-1,category:"PATIENCE",description:"Opened coffee shop"}
            action.push(JSON.stringify(action2))
        } else {
            let action1 = {score:-2,category:"MONEY",description:"Bought stock for shop"}
            action.push(JSON.stringify(action1))
		  }
    }
    if(shopOpened) {
		if(day > currentDay) { // a new day, decrement counter
			currentDay = day
			shopDaycount++
		}
		//50/50 chance of having a coffee shop customer
		if((throws[1] % 2) == 0) {
			   if(throws[2] < 3) {
               let action4 = {score:2,category:"MONEY",description:"Sold a coffee"}
               action.push(JSON.stringify(action4))
				}
			   if(throws[2] > 2 && throws[2] < 5) {
               let action4 = {score:1,category:"MONEY",description:"Sold a tea"}
               action.push(JSON.stringify(action4))
				}
			   if((throws[1] > 3) && throws[2] > 4) {
               let action4 = {score:3,category:"MONEY",description:"Sold a drink and a cake/sandwich"}
               action.push(JSON.stringify(action4))
				}
		} else {
			if(throws[2] >= 5) {
            let action3 = {score:1,category:"PATIENCE",description:"Close coffee shop"}
            action.push(JSON.stringify(action3))
			}
		}
	}
    if(verbose) {
        var m = ""
        if(shopOpened) {
            m = `, on day ${shopDaycount}`
        }
		console.log(`you threw ${throws[0]},${throws[1]}. coffee shop is ${runStatus}${m}`)
	}
    return action;
}
