class QuestionController {
    questions
    observers

    constructor(questions, template) {
        this.questions = questions
        this.observers = template
    }

    notifyAndAddQuestion(newQuestion, ind){
        this.observers.innerHTML += newQuestion.draw({...newQuestion},ind)
    }

    addAQuestion(){
        let newQuestion = dummyQuestion('single')
        this.questions.push(newQuestion)
        this.notifyAndAddQuestion(newQuestion, this.questions.length-1)
    }



    outputAll(container){
        container.innerHTML=''
        for (let i = 0 ; i < this.questions.length; i++){
           container.innerHTML+=this.questions[i].draw({...this.questions[i]}, i)
        }
        assignListeners()

    }
}
function assignListeners(){
    document.querySelectorAll('.question-answers-item').forEach((item,ind)=>{
        item.addEventListener('input',(e)=>{
            change(e.target)
        })
    })
    document.querySelectorAll('.question-types').forEach((item,ind)=>{
        item.addEventListener('change', (e)=>{
            console.log(e.target.value)
            changeTypeOfQuestion(ind, e.target.value)
            item.textContent = e.target.value
        })
    })
    document.querySelectorAll('.input-radio').forEach((item,ind)=>{
        item.addEventListener('change', (e)=>{
            changeCorrectAnswers(e.target.name, e.target.value)
        })
    })

    document.querySelectorAll('.input-radio-MCQ').forEach((item,ind)=>{
        item.addEventListener('change', (e)=>{
            changeMcqCorrectAnswers(e.target.name, e.target.value)
        })
    })
}

function dummyQuestion(type){
    switch (type){
        case 'single':
            return new SingleAnswerQuestion("dummy description",
                {A:{value:1,correct:false}, B:{value:1,correct:true},C:{value:1,correct:false},D:{value:1,correct:false}}
                , 'B')
        case 'MCQ':
            return new MCQQuestion("dummy description", {A:{value:1,correct:false}, B:{value:1,correct:false},C:{value:1,correct:false},D:{value:1,correct:false}}, [])
        case 'boolean':
            return new BooleanQuestion("dummy description", {A:{value:true,correct:true}, B:{value:false,correct:false}}, 'D')
    }

}