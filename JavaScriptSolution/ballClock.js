//Ball class

function Ball(id) {
    this.id = id;
}

//BallDisplay class (inherits from Ball)

BallDisplay.prototype = Object.create(Ball.prototype);
BallDisplay.prototype.constructor = BallDisplay;

function BallDisplay(id) {
    Ball.call(this, id);
}

//Clock class
function Clock(numBalls, mode) {
    this._minutes = [];
    this._fiveMinutes = [];
    this._hours = [];
    this._queue = [];
    this._numBalls = numBalls;
    var iterator = 0;
    var length_amount = numBalls;
    if (mode && mode === 1) {
    	iterator = 0;
    	length_amount = numBalls;
    } else if (mode && mode === 2) {
    	iterator = 1;
    	length_amount = numBalls + 1;
    }
    for (var i = iterator; i < length_amount; i++) {
        var ball = new BallDisplay(i);
        this._queue.push(ball);
    }
}

Clock.prototype.increment = function() {
    var ball = this._queue.shift();
    if (this._minutes.length < 4) {
        this._minutes.push(ball);
    } else {
        while (this._minutes.length) {
            this._queue.push(this._minutes.pop());
        }

        if (this._fiveMinutes.length < 11) {
            this._fiveMinutes.push(ball);
        } else {
            while (this._fiveMinutes.length) {
                this._queue.push(this._fiveMinutes.pop());
            }

            if (this._hours.length < 11) {
                this._hours.push(ball);
            } else {
                while (this._hours.length) {
                    this._queue.push(this._hours.pop());
                }
                this._queue.push(ball);
            }
        }
    }
}

Clock.prototype.home = function() {
    if (this._queue.length != this._numBalls)
        return false;
    for (var i = 0; i < this._numBalls; i++) {
        if (this._queue[i].id != i)
            return false;
    }
    return true;
}

//create instances/page functionality
var days = [];

function findDays(balls, mode=1, passedMinutes=0) {
    var count=days[balls];
    if (mode && mode === 1) {
    	if(!count) {
		count=0;
		var clock=new Clock(balls, mode);
		do {
			clock.increment();
			count++;
		}
		while(!clock.home())
		days[balls]=count;
		}
    }
    if (mode && mode === 2) {
    	var clock=new Clock(balls, mode);
    	for (var i = 0; i < passedMinutes; i++) {
    		clock.increment();
    	}
    	console.log(clock._minutes);
    	console.log(clock._fiveMinutes);
    	console.log(clock._hours);
    	console.log(clock._queue);
    }
    return count / 60 / 24;
}

var balls = 30;
var count = findDays(balls, 1, 325);
console.log(balls + " balls cycle after " + count + " days.");