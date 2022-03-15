var saleStarted = false
var saleDaycount = 0
var currentDay = 0
var verbose = false

function setVerbose(v) {
    verbose = v
}

function getName() {
    return "simple, sale expansion"
}

function getRequiredThrows() {
    return 2
}

function run(day, throws) {
    var req = getRequiredThrows()
    if(throws.length != req) {
        throw `you must provide ${req} dice throws.`
    }
    var runStatus = ""
    if(saleStarted && saleDaycount > 0) {
        runStatus = "running"
    }
    var action = []
    //1 in 3 chance that sale starts if not already started
    if(throws[0] < 3) {
        if(!saleStarted) {
            saleStarted = true
            saleDaycount = 6
            currentDay = day
            runStatus = "starting"
            let action1 = {score:-5,category:"MONEY",description:"Bought more books for sale"}
            action.push(JSON.stringify(action1))
            let action2 = {score:-1,category:"PATIENCE",description:"Start book sale"}
            action.push(JSON.stringify(action2))
        }
    }
    if(saleStarted) {
		if(day > currentDay) { // a new day, decrement counter
			currentDay = day
			saleDaycount--
			if(saleDaycount <= 0) {
				saleStarted = false
                let action3 = {score:1,category:"PATIENCE",description:"End book sale"}
                action.push(JSON.stringify(action3))
			}
		}
		//50/50 chance of selling a book during sale
		if((throws[1] % 2) == 0) {
            let action4 = {score:1,category:"MONEY",description:"Sold a book in the sale"}
            action.push(JSON.stringify(action4))
		}
	}
    if(verbose) {
        var m = ""
        if(saleStarted) {
            m = `, on day ${15-saleDaycount}`
        }
		console.log(`you threw ${throws[0]},${throws[1]}. sale is ${runStatus}${m}`)
	}
    return action;
}