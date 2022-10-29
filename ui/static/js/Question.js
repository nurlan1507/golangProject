
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
        this.notifyAndAddQuestion(newQuestion,this.questions.length-1)
    }

    changeTypeOfQuestion(ind,type){
        this.questions[ind] = dummyQuestion(type)
        console.log({...this.questions[ind]})
        this.observers.innerHTML =''
        this.outputAll(this.observers)
        document.querySelectorAll('.question-answers-item').forEach((item,ind)=>{
            item.addEventListener('input',(e)=>{
                change(e.target)
            })
        })
    }

    getLen(){
        console.log(this.questions.length)
    }
    outputAll(container){
        container.innerHTML=''
        this.getLen()
        for (let i = 0 ; i < this.questions.length; i++){
           container.innerHTML+=this.questions[i].draw({...this.questions[i]}, i)
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