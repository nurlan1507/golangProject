class PaginationTest{
    questions
    page=1
    constructor(questions) {
        this.questions = questions
    }
    definePagesAmount(){
        return Math.ceil(this.questions.length/5)
    }
    drawPag(){
        for (let i = 0 ; i < this.definePagesAmount(); i++){
            document.querySelector('.pagination_section').innerHTML+=
                `<a href="/createTest?page=${++i}" title="Algorithm">${++i}</a> `
        }
    }


}