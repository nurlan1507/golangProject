
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
        this.questions.push({...newQuestion})
        this.notifyAndAddQuestion(newQuestion, this.questions.length-1)
    }


    outputAll(container){
        container.innerHTML=''
        for (let i = 0 ; i < this.questions.length; i++){
           container.innerHTML+=this.questions[i].draw()
            }
    }
}

function dummyQuestion(type){
    switch (type){
        case 'single':
            return new SingleAnswerQuestion("dummy description", {A:1, B:2,C:3,D:4}, 'D')
        case 'MCQ':
            return new MCQQuestion("dummy description", {A:1, B:2,C:3,D:4}, 'D')
        case 'boolean':
            return new BooleanQuestion("dummy description", {A:true, B:false}, 'D')
    }

}