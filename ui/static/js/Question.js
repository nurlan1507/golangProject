// var questionTemplates =[
//     {
//         type:'single',
//         description:"What is 10/2?",
//         answers: {
//             A: '3',
//             B: '5',
//             C: '115',
//             D: '22'
//         },
//         correctAnswer: 'A'
//     },
//     {
//         type:'boolean',
//         description: "you gay?",
//         answers: {
//             A:true,
//             B:false,
//         },
//         correctAnswer: 'A'
//     },
//     {
//         type:'MCQ',
//         description:"What is 10/2?",
//         answers: {
//             A: '5.0',
//             B: '5',
//             C: '115',
//             D:'4',
//         },
//         correctAnswer: ['A', 'B']
//     },
// ]
//
// class Question {
//     questions
//
//     constructor(questions) {
//         this.questions = questions
//     }
//
//     outputSingleAnswer(question,ind) {
//         var output = ''
//         var answers = []
//         for (let key in question.answers) {
//                 answers.push(
//                     '<div style="display: flex">'+
//                     '<input type="radio" style="width: 20px; height: 20px" name="question'+`${ind}` +'" value="'+key+ind+'">'
//                     +'<div >'  +key+':'+ '<span contenteditable="true" id='+`${key}`+ ind+' class="question-answers-item" onchange= ' + `change(event.target)`+
//                     '>' +'  '+ question.answers[key]+'</span> </div> '+
//                     '<button type="submit" class="add-question-btn" formaction="/addQuestion">add</button>'+
//                     '</div>'
//                 )
//         }
//         output = '<div class="question">' +
//             '<select class="question-types">' +
//             '<option value="1">One answer</option>' +
//             '<option value="2">MCQ</option>' +
//             '<option value="3">Boolean</option>' +
//             '</select>' +
//             '<div class="question-name" contentEditable="true">' + question.description + '</div>' +
//             '<div class="question-answers">' + answers.join("") + '</div>'
//             + '</div>'
//         return output
//     }
//
//     outputMCQQuestion(question,ind) {
//         var output = ''
//         var answers = []
//         for (let key in question.answers) {
//             answers.push(
//                 '<div style="display: flex">'+
//                 '<input type="checkbox" style="width: 20px; height: 20px" name="question'+`${ind}` +'" value="'+key+ind+'">'
//                 +'<div >'  +key+':'+ '<span contenteditable="true" id='+`${key}`+ ind+' class="question-answers-item" onchange= ' + `change(event.target)`+
//                 '>' +'  '+ question.answers[key]+'</span> </div> '+
//                 '<button type="submit" class="add-question-btn" formaction="/addQuestion">add</button>'+
//                 '</div>'
//             )
//         }
//         output = '<div class="question">' +
//             '<select class="question-types">' +
//             '<option value="1">One answer</option>' +
//             '<option value="2">MCQ</option>' +
//             '<option value="3">Boolean</option>' +
//             '</select>' +
//             '<div class="question-name" contentEditable="true">' + question.description + '</div>' +
//             '<div class="question-answers">' + answers.join("") + '</div>'
//             + '</div>'
//         return output
//     }
//
//     outputBoolean(question,ind) {
//         var output = ''
//         var answers = []
//         for (let key in question.answers) {
//             answers.push(
//                 '<div style="display: flex">' +
//                 '<input type="radio"  style= "width:20px;height: 20px" name="question" value="' + key + '">'
//                 + '<div >' + key + ':' + '<span contenteditable="true" id=' + `${key}`+ind+' >' + '  ' + question.answers[key] + '</span> </div> </div>'
//             )
//         }
//         output = '<div class="question">' +
//             '<select class="question-types">' +
//             '<option value="1">One answer</option>' +
//             '<option value="2">MCQ</option>' +
//             '<option value="3">Boolean</option>' +
//             '</select>' +
//             '<div class="question-name" contentEditable="true">' + question.description + '</div>' +
//             '<div class="question-answers">' + answers.join("") + '</div>'
//             + '</div>'
//         return output
//     }
//
//     outputAll(container){
//         for (let i = 0 ; i < this.questions.length; i++){
//             switch (this.questions[i].type){
//                 case "MCQ":
//                     container.innerHTML+=this.outputMCQQuestion(this.questions[i],i)
//                     break
//                 case "boolean":
//                     container.innerHTML+=this.outputBoolean(this.questions[i],i)
//                     break
//                 case "single":
//                     container.innerHTML+=this.outputSingleAnswer(this.questions[i],i)
//                     break
//             }
//         }
//     }
// }
