var questionTemplates =[
    {
        type:'single',
        description:"What is 10/2?",
        answers: {
            A: '3',
            B: '5',
            C: '115',
            D: '22'
        },
        correctAnswer: 'A'
    },
    {
        type:'boolean',
        description: "you gay?",
        answers: {
            A:true,
            B:false,
        },
        correctAnswer: 'A'
    },
    {
        type:'MCQ',
        description:"What is 10/2?",
        answers: {
            A: '5.0',
            B: '5',
            C: '115',
            D:'4',
        },
        correctAnswer: ['A', 'B']
    },
]

class Question{
    questions
    constructor(questions ) {
        this.questions = questions
    }

    outputSingleAnswer(question){
        var output =''
        var answers = []
        for (let key in question.answers){
            answers.push(
                '<div style="display: flex">'+
                '<input type="radio" style="width: 20px; height: 20px" name="question" value="'+key+'">'
                +'<div >'  +key+':'+ '<span contenteditable="true">' +'  '+ question.answers[key]+'</span> </div> </div>'
            )
        }
        output = '<div class="question">'+
            '<select class="question-types">'+
            '<option value="1">One answer</option>'+
            '<option value="2">MCQ</option>'+
            '<option value="3">Boolean</option>'+
            '</select>'+
            '<div class="question-name" contentEditable="true">' + question.description+ '</div>' +
            '<div class="question-answers">'+answers.join("")+'</div>'
            +'</div>'
        return output
    }

    outputMCQQuestion(question){
        var output =''
        var answers = []
        for (let key in question.answers){
            answers.push(
                '<div style="display: flex">'+
                '<input type="checkbox" style="width: 20px; height: 20px" name="question" value="'+key+'">'
                +'<div >'  +key+':'+ '<span contenteditable="true">' +'  '+ question.answers[key]+'</span> </div> </div>'
            )
        }
        output = '<div class="question">'+
            '<select class="question-types">'+
            '<option value="1">One answer</option>'+
            '<option value="2">MCQ</option>'+
            '<option value="3">Boolean</option>'+
            '</select>'+
            '<div class="question-name" contentEditable="true">' + question.description+ '</div>' +
            '<div class="question-answers">'+answers.join("")+'</div>'
            +'</div>'
        return output
    }

    outputBoolean(question){
        var output = ''
        var answers = []
        for (let key in question.answers){
            answers.push(
                '<div style="display: flex">'+
                '<input type="radio" style="width: 20px; height: 20px" name="question" value="'+key+'">'
                +'<div >'  +key+':'+ '<span contenteditable="true">' +'  '+ question.answers[key]+'</span> </div> </div>'
            )
        }
        output = '<div class="question">'+
            '<select class="question-types">'+
            '<option value="1">One answer</option>'+
            '<option value="2">MCQ</option>'+
            '<option value="3">Boolean</option>'+
            '</select>'+
            '<div class="question-name" contentEditable="true">' + question.description+ '</div>' +
            '<div class="question-answers">'+answers.join("")+'</div>'
            +'</div>'
        return output
    }
}