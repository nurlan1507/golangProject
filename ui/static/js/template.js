class Question{
    description
    answers
    correctAnswer
    constructor( description, answers, correctAnswer) {
        this.description=description
        this.answers=answers
        this.correctAnswer=correctAnswer
    }
    draw(question,ind){

    }
}

class SingleAnswerQuestion extends Question{
    type = 'single'

    constructor(description,answers,correctAnswer) {
        super(description,answers ,correctAnswer);
    }
    draw(question,ind) {
        var output = ''
        var answers = []
        for (let key in question.answers) {
            answers.push(
                `<div style="display: flex">
                <input type="radio" class="input-radio" style="width: 20px; height:20px" name="${ind}" value="${key}">
                <div>${key} :<span contenteditable="true" id="${key}${ind}" class="question-answers-item" onchange=change(e.target)>
                ${question.answers[key]}</span></div></div>`
            )
        }
        output = `<div class="question">
        <select class="question-types">
            <option value="single" selected>single</option>
            <option value="MCQ">MCQ</option>
            <option value="boolean">Boolean</option>
        </select>
        <div class="question-name" contentEditable="true">${question.description}</div>
        <div class="question-answers">${answers.join("")}</div>
        </div>`
        return output

    }

    outputSingleAnswer(question,ind) {

    }

}

class BooleanQuestion extends Question{
    type = 'boolean'
    constructor(description, answers, correctAnswer) {
        super(description,answers,correctAnswer)
    }
    draw(question,ind) {
        var output = ''
        var answers = []
        for (let key in question.answers) {
            answers.push(
                `<div style="display:flex;">
                <input type="radio" class="input-radio" style="width: 20px; height: 20px" name="${ind}" value="${key}">
                <div>${key}: <span>${question.answers[key]}</span></div></div>`
            )
        }
        output = `<div class="question">
        <select class="question-types">
            <option value="single">single</option>
            <option value="MCQ">MCQ</option>
            <option value="boolean">Boolean</option>
        </select>
        <div class="question-name" contentEditable="true">${question.description}</div>
        <div class="question-answers">${answers.join("")}</div>
        </div>`
        return output
    }

    outputBoolean(question,ind) {

    }
}

class MCQQuestion extends Question{
    type='MCQ'

    constructor( description, answers, correctAnswer) {
        super(description,answers,correctAnswer)
    }
    draw(question,ind) {
        var output = ''
        var answers = []
        for (let key in question.answers) {
            answers.push(
                `<div style="display: flex">
                <input type="checkbox" class="input-radio-MCQ"  style="width: 20px; height:20px" name="${ind}" value="${key}">
                <div>${key} :<span contenteditable="true" id="${key}${ind}" class="question-answers-item">
                ${question.answers[key]}</span></div></div>`
            )
        }
        output =`<div class="question">
        <select class="question-types">
            <option value="single">single</option>
            <option value="MCQ">MCQ</option>
            <option value="boolean">Boolean</option>
        </select>
        <div class="question-name" contentEditable="true">${question.description}</div>
        <div class="question-answers">${answers.join("")}</div>
        </div>`
        return output
    }

    outputMCQQuestion(question,ind) {

    }
}
