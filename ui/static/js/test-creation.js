// import {Question} from "/static/js/Question.js";

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





function change(eTarget){
    console.log(eTarget.id)
    questions.questions[eTarget.id[1]].answers[eTarget.id[0]] = eTarget.textContent
}
function changeTypeOfQuestion(ind,type){
    questions.questions[ind] =  dummyQuestion(type)

}

let questions = new QuestionController([], document.querySelector('.question-holder'))
document.getElementById('addNewQuestion').addEventListener("click", ()=>{
    console.log('click')
    questions.addAQuestion()
    document.querySelectorAll('.question-answers-item').forEach((item,ind)=>{
        item.addEventListener('input',(e)=>{
            change(e.target)

        })
    })
    document.querySelectorAll('.question-types').forEach((item,ind)=>{
        item.addEventListener('change', (e)=>{
            console.log(e.target.value)

        })
    })
})




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





