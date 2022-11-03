
function change(eTarget){
    console.log(eTarget.id)
    questions.questions[eTarget.id[1]].answers[eTarget.id[0]].value = eTarget.textContent
}
function changeTypeOfQuestion(ind,type){
    questions.questions[ind] =  dummyQuestion(type)
    questions.observers.innerHTML = ''
    questions.outputAll(document.querySelector('.question-holder'))
}
function changeCorrectAnswers(ind,key){
    questions.questions[ind].answers[questions.questions[ind].correctAnswer].correct=false
    questions.questions[ind].answers[key].correct = true
    questions.questions[ind].correctAnswer = key
}
function changeMcqCorrectAnswers(ind,key){
    let ex = false
    console.log( questions.questions[ind].correctAnswer)
    questions.questions[ind].correctAnswer.forEach((item,it)=>{
        if (item === key){
            questions.questions[ind].correctAnswer.splice(it,1)
            ex = true
        }
    });
    questions.questions[ind].answers[key].correct =   questions.questions[ind].answers[key].correct!==true
    if (ex ===true){
        return
    }else{
        questions.questions[ind].correctAnswer.push(key)
        return;
    }
}

let questions = new QuestionController([], document.querySelector('.question-holder'))
document.getElementById('addNewQuestion').addEventListener("click", ()=>{
    console.log('click')
    questions.addAQuestion()
    assignListeners()
})




document.getElementById('create_test').addEventListener("click",()=>{
    console.log(questions.questions)
    var jsonData ={}
    jsonData.questions = 23
    console.log(jsonData)
    var fetchOptions = {
        method: "POST",
        header: new Headers({
            'Content-Type': "application/json",
        }),
        body: JSON.stringify(jsonData)
    }
    const url ='http://localhost:4000/createTest'
    fetch(url,fetchOptions).then((res)=>{
        return res.json()
    }).then((res)=>{
        console.log(res.status)
    })
})






