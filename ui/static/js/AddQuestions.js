const AddQuestions =async(data)=>{
    var jsonData ={}
    jsonData.questions = data
    var fetchOptions = {
        method: "POST",
        header: new Headers({
            'Content-Type': "application/json",
        }),
        body: JSON.stringify(jsonData)
    }
    try{
        const url ='http://localhost:4000/createTest'
        const res = await fetch(url,fetchOptions)
    }catch (e) {
        console.log(e)
    }


}