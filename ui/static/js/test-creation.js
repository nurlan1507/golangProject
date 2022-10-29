// import {Question} from "/static/js/Question.js";
function change(eTarget){
    console.log(eTarget.id)
    myQuestion[eTarget.id[1]].answers[eTarget.id[0]] = eTarget.textContent
}
var myQuestion = [
    {
        type:"single",
        description:"What is 10/2?",
        answers: {
            A: '3',
            B: '5',
            C: '115',
            D:'2'
        },
        correctAnswer: 'A'
    },
    {
        type:"MCQ",
        description:"What is 10/2?",
        answers: {
            A: '3',
            B: '5',
            C: '115',
            D:'2'
        },
        correctAnswer: 'B'
    }
]

class Question {
    questions
    constructor(questions) {
        this.questions = questions
    }

    outputSingleAnswer(question,ind) {
        var output = ''
        var answers = []
        for (let key in question.answers) {
            answers.push(
                `<div style="display: flex">
                <input type="radio" style="width: 20px; height:20px" name="question+${ind}" value="${key}">
                <div>${key} :<span contenteditable="true" id="${key}${ind}" class="question-answers-item" onchange=change(e.target)>
                ${question.answers[key]}</span></div></div>`
            )
        }
        output = '<div class="question">' +
            '<select class="question-types">' +
            '<option value="1">One answer</option>' +
            '<option value="2">MCQ</option>' +
            '<option value="3">Boolean</option>' +
            '</select>' +
            '<div class="question-name" contentEditable="true">' + question.description + '</div>' +
            '<div class="question-answers">' + answers.join("") + '</div>'
            + '</div>'
        return output
    }

    outputMCQQuestion(question,ind) {
        var output = ''
        var answers = []
        for (let key in question.answers) {
            answers.push(
                `<div style="display: flex">
                <input type="checkbox"  style="width: 20px; height:20px" name="question+${ind}" value="${key}">
                <div>${key} :<span contenteditable="true" id="${key}${ind}" class="question-answers-item" oninput="change(e.target)">
                ${question.answers[key]}</span></div></div>`
            )
        }
        output = '<div class="question">' +
            '<select class="question-types">' +
            '<option value="1">One answer</option>' +
            '<option value="2">MCQ</option>' +
            '<option value="3">Boolean</option>' +
            '</select>' +
            '<div class="question-name" contentEditable="true">' + question.description + '</div>' +
            '<div class="question-answers">' + answers.join("") + '</div>'
            + '</div>'
        return output
    }

    outputBoolean(question,ind) {
        var output = ''
        var answers = []
        for (let key in question.answers) {
            answers.push(
                `<div style="display:flex;">
                <input type="radio" style="width: 20px; height: 20px" name="question+${ind}" value="${question.answers[key]}">
                <div>${key}: <span>${question.answers[key]}</span></div></div>`
            )
        }
        output = '<div class="question">' +
            '<select class="question-types">' +
            '<option value="1">One answer</option>' +
            '<option value="2">MCQ</option>' +
            '<option value="3">Boolean</option>' +
            '</select>' +
            '<div class="question-name" contentEditable="true">' + question.description + '</div>' +
            '<div class="question-answers">' + answers.join("") + '</div>'
            + '</div>'
        return output
    }

    outputAll(container){
        for (let i = 0 ; i < this.questions.length; i++){
            switch (this.questions[i].type){
                case "MCQ":
                    container.innerHTML+=this.outputMCQQuestion(this.questions[i],i)
                    break
                case "boolean":
                    container.innerHTML+=this.outputBoolean(this.questions[i],i)
                    break
                case "single":
                    container.innerHTML+=this.outputSingleAnswer(this.questions[i],i)
                    break
            }
        }
    }
}




let questions = new Question(myQuestion)
questions.outputAll(document.querySelector('.question-holder'))

//
// function generateQuestions(container, questionArr){
//     let output = [];
//     for (let l = 0; l < questionArr.length ; l++){ //2
//         let answers = [];
//         for (var key in questionArr[l].answers){                 //key is A B C
//             console.log(l)
//             answers.push(
//                 '<div style="display: flex">'+
//                 '<input type="radio" style="width: 20px; height: 20px" name="question'+`${l}` +'" value="'+key+l+'">'
//                 +'<div >'  +key+':'+ '<span contenteditable="true" id='+`${key}`+ l+' class="question-answers-item" onchange= ' + `change(event.target)`+
//                 '>' +'  '+ questionArr[`${l}`].answers[key]+'</span> </div> '+
//                 '<button type="submit" class="add-question-btn" formaction="/addQuestion">add</button>'+
//                 '</div>'
//             )
//         }
//         output.push(
//             '<div class="question">'+
//             '<select class="question-types">'+
//                 '<option value="1">One answer</option>'+
//                '<option value="2">MCQ</option>'+
//                '<option value="3">Boolean</option>'+
//             '</select>'+
//             '<div class="question-name" contentEditable="true">' + questionArr[l].description+ '</div>' +
//             '<div class="question-answers">'+answers.join("")+'</div>'
//             +'</div>'
//         )
//     }
//     container.innerHTML = output.join("")
// }
//
// function generateNewQuestionTemplate(container, questionArr){
//     var newQuestion = {
//             description:"What is 10/2?",
//             answers: {
//                 A: '3',
//                 B: '5',
//                 C: '115'
//             },
//             correctAnswer: 'A'
//         };
//
//     container.innerHTML +=  '<div class="question">'+
//         '<select class="question-types">'+
//         '<option value="1">One answer</option>'+
//         '<option value="2">MCQ</option>'+
//         '<option value="3">Boolean</option>'+
//         '</select>'+
//         '<div class="question-name" contentEditable="true">' + newQuestion.description+ '</div>' +
//         '<div class="question-answers">'+ + '</div>'
//         +'</div>'
//
// }
//
// generateQuestions(document.querySelector('.question-holder'), myQuestion)



document.querySelectorAll('.question-answers-item').forEach((item,ind)=>{
    item.addEventListener('input',(e)=>{
        change(e.target)
    })
})

