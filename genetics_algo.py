#!/usr/bin/env python
# coding: utf-8

get_ipython().run_line_magic('matplotlib', 'inline')
import numpy as np, random, operator, pandas as pd, matplotlib.pyplot as plt
import math

componentsKey = 'Components'
machinesKey = 'Machines'
processNameKey = 'processName'
processTimeKey = 'processTime'
machineNameKey = 'machineName'
componentNameKey = 'componentName'
timeKey = 'timeStartEnd'
durationKey = 'duration'
componentSetKey = 'setKey'
cycleTimeKey = 'cycleTimeKey'

## Change this !
def constructDict(arrayClass):
    dict = {}
    for obj in arrayClass:
        dict[obj.getName()] = obj
    return dict

class PlanCalculator:
    def __init__(self, constraints, plan):
        self.machines = constraints[machinesKey]
        self.components = constraints[componentsKey]
        self.plan = plan
        # {'M1':[{componentNameKey: 'C1', timeKey: [0,3]}]}
        self.greedyArrangement = {}
        self.bestArrangement = {}
        # [{setKey:set(), cycleTime: int}]
        self.cycleTimeComponentSet = []
    def getMachinesNameUsedByPlanIndex(self, planIndex):
        componentPlan = self.plan[planIndex]
        machineUsedNames = set()
        for machinePlan in componentPlan:
            machineUsedNames.add(machinePlan.getName())
        return machineUsedNames
    def getMachinesNameUsed(self, componentName):
        for i, component in enumerate(self.components):
            if component.getName() == componentName:
                return self.getMachinesNameUsedByPlanIndex(i)
        return set()
    def getCycleTimeFromExistingSet(self, componentName):
        for setObj in self.cycleTimeComponentSet:
            if componentName in setObj[componentSetKey]:
                return setObj[cycleTimeKey], True
        return 0, False
    def buildCycleTime(self, planIndex):
        machineUsedNames = self.getMachinesNameUsedByPlanIndex(planIndex)
        queueSet = machineUsedNames
        componentsName = set([self.components[planIndex].getName()])
        while True:
            newQueueSet = set()
            for machineName in queueSet:
                componentScheds = self.bestArrangement[machineName]
                newComponentsName = set()
                for componentSched in componentScheds:
                    componentName = componentSched[componentNameKey]
                    if componentName not in componentsName:
                        componentsName.add(componentName)
                        newComponentsName.add(componentName)
                for newComponentName in newComponentsName:
                    machinesNameSet = self.getMachinesNameUsed(newComponentName)
                    for machinesName in machinesNameSet:
                        if machinesName not in machineUsedNames:
                            newQueueSet.add(machinesName)
            if len(newQueueSet) == 0:
                break
            else:
                for newName in newQueueSet:
                    machineUsedNames.add(newName)
                queueSet = newQueueSet

        maxTime = 1
        for machineName in machineUsedNames:
            machineSchedule = self.bestArrangement[machineName]
            machineCycleTime = machineSchedule[len(machineSchedule)-1][timeKey][1] - machineSchedule[0][timeKey][0]
            maxTime = max(machineCycleTime, maxTime)
        return maxTime, machineUsedNames
    def addCycleTimeSet(self, setObj):
        self.cycleTimeComponentSet.append(setObj)

    def getCycleTime(self, planIndex):
        componentName = self.components[planIndex].getName()
        cycleTime, found = self.getCycleTimeFromExistingSet(componentName)
        if not found:
            cycleTime, machineUsedNames = self.buildCycleTime(planIndex)
            self.addCycleTimeSet({cycleTimeKey: cycleTime, componentSetKey: machineUsedNames})
        return cycleTime

    def getBestArrangement(self):
        return self.bestArrangement
    def calculateGreedy(self):
        # queue item : {"machineName":string, "componentName":string}
        # TODO: the limitation is that it use queue, so first component has higher priority. Improve this
        queue = []
        i = 0
        # transform all component plan into queue
        while True:
            count = 0
            for index, componentPlan in enumerate(self.plan):
                if i >= len(componentPlan):
                    count +=1
                    continue
                machineChosen = componentPlan[i]
                queue.append({machineNameKey: machineChosen.getName(), componentNameKey: self.components[index].getName(),
                              durationKey:self.components[index].getProcessTime(i)})
            i += 1
            if count == len(self.plan):
                break
        # assign machine task based on queue
        componentTaskTime = {}
        for index, queueTask in enumerate(queue):
            queueTaskMachineName = queueTask[machineNameKey]
            queueTaskComponentName = queueTask[componentNameKey]
            duration = queueTask[durationKey]
            if queueTaskMachineName not in self.greedyArrangement:
                startTime = 0
                if queueTaskComponentName in componentTaskTime:
                    startTime = componentTaskTime[queueTaskComponentName]
                self.greedyArrangement[queueTaskMachineName] =[{componentNameKey:queueTaskComponentName, timeKey:[startTime, startTime + duration]}]
                componentTaskTime[queueTaskComponentName] = startTime + duration
            else:
                startTime = self.greedyArrangement[queueTaskMachineName][len(self.greedyArrangement[queueTaskMachineName])-1][timeKey][1]
                if queueTaskComponentName in componentTaskTime:
                    startTime = max(componentTaskTime[queueTaskComponentName], startTime)
                self.greedyArrangement[queueTaskMachineName].append({componentNameKey:queueTaskComponentName, timeKey:[startTime, startTime + duration]})
                componentTaskTime[queueTaskComponentName] = startTime + duration
    def calculate(self):
        self.calculateGreedy()
        self.bestArrangement = self.greedyArrangement


class Machine:
    def __init__(self, name, cost, processNames):
        self.name = name
        self.cost = cost
        self.processNames = processNames
    def hasProcess(self, processName):
        return processName in self.processNames
    def getCost(self):
        return self.cost
    def getName(self):
        return self.name

class Component:
    def __init__(self, name, price, processes):
        self.name = name
        self.price = price
        self.processes = processes

    def getRandomPlan(self, machines):
        # [[M1],[],[M1,M3]...] means C1 has M1, C2 is not produced, C3 has process M1,M3
        plan = []
        for process in self.processes:
            name = process[processNameKey]
            candidateMachines = []
            for machine in machines:
                if machine.hasProcess(name):
                    candidateMachines.append(machine)
            if len(candidateMachines) == 0:
                print("there is no process to create component "+ self.name)
                continue
            pickedMachine = candidateMachines[random.randint(0, len(candidateMachines)-1)]
            plan.append(pickedMachine)
        return plan

    def getProfit(self, processPlanMachines):
        cost = 0
        for i in range(0,len(processPlanMachines)):
            cost += processPlanMachines[i].getCost() * self.processes[i][processTimeKey]
        return self.price - cost
    def getProcessName(self,index):
        return self.processes[index][processNameKey]
    def getProcessTime(self, processIndex):
        return self.processes[processIndex][processTimeKey]
    def getName(self):
        return self.name

def calculateProfit(constraints, plan):
    # calculate the profit and count type of product produced (min N to be counted as 1 otherwise its 0)
    calculator = PlanCalculator(constraints, plan)
    calculator.calculate()
    components = constraints[componentsKey]
    profit = 0
    productTypeCount = 0
    for i in range(0, len(plan)):
        if len(plan[i]) > 0:
            productTypeCount += 1
            profit += components[i].getProfit(plan[i]) / calculator.getCycleTime(i)

    return profit, productTypeCount

def normalizePopulationScore(constraints, population):
    # normalize the distribution, can do this by simply get  the avg value
    profitPopSum = 0
    productSum = 0
    profits = []
    productSums = []
    for plan in population:
        profit, numberOfProduct = calculateProfit(constraints, plan)
        profitPopSum += profit
        productSum += numberOfProduct
        profits.append(profit)
        productSums.append(numberOfProduct)

    length = len(population)
    return profitPopSum/length, productSum/length, profits, productSums

def calculateScore(profit, numberOfProduct, avgProfit, avgProduct):
    # finalScore = (c1 * profit / avgProfit) + (c2 * numberOfProduct/ avgProduct)
    finalScore = (c1 * profit) + (c2 * numberOfProduct)
    return finalScore

def createRandomPlan(constraints):
    # create random plan here
    # define Plan as array of components order of machining
    # [[M1],[],[M1,M3]...] means C1 has M1, C2 is not produced, C3 has process M1,M3
    components = constraints[componentsKey]
    machines = constraints[machinesKey]
    arrayPlan = []
    for component in components:
        # decide whether to produce or not
        if random.random() < 0.5:
            arrayPlan.append([])
        else:
            arrayPlan.append(component.getRandomPlan(machines))
    return arrayPlan

def breed(parent1, parent2):
    # crossover between 2 plan to get a new plan
    # TODO: improve ?
    r = random.randint(0, len(parent1)-1)
    newPlan = []
    for i in range(0, len(parent1)):
        if i < r:
            newPlan.append(parent1[i])
        else:
            newPlan.append(parent2[i])
    return newPlan

def mutate(constraints, plan, mutationRate):
    components = constraints[componentsKey]
    machines = constraints[machinesKey]
    componentMutationRate = mutationRate/ len(plan)
    for i in range(0, len(plan)):
        if(random.random() < componentMutationRate):
            plan[i] = components[i].getRandomPlan(machines)
    return plan


## No need to change

def initialPopulation(popSize, constraints):
    population = []
    for i in range(0, popSize):
        population.append(createRandomPlan(constraints))
    return population

def rankPop(constraints, population):
    fitnessResults = {}
    avgProfit, avgProduct, profits, productSums = normalizePopulationScore(constraints, population)
    for i in range(0,len(population)):
        fitnessResults[i] = calculateScore(profits[i], productSums[i], avgProfit, avgProduct)
    return sorted(fitnessResults.items(), key = operator.itemgetter(1), reverse = True)

def selection(popRanked, eliteSize):
    selectionResults = []
    df = pd.DataFrame(np.array(popRanked), columns=["Index","Fitness"])
    df['cum_sum'] = df.Fitness.cumsum()
    df['cum_perc'] = 100*df.cum_sum/df.Fitness.sum()
    
    for i in range(0, eliteSize):
        selectionResults.append(popRanked[i][0])
    for i in range(0, len(popRanked) - eliteSize):
        pick = 100*random.random()
        for i in range(0, len(popRanked)):
            if pick <= df.iat[i,3]:
                selectionResults.append(popRanked[i][0])
                break
    return selectionResults

def matingPool(population, selectionResults):
    matingpool = []
    for i in range(0, len(selectionResults)):
        index = selectionResults[i]
        matingpool.append(population[index])
    return matingpool

def breedPopulation(matingpool, eliteSize):
    children = []
    length = len(matingpool) - eliteSize
    pool = random.sample(matingpool, len(matingpool))

    for i in range(0,eliteSize):
        children.append(matingpool[i])
    
    for i in range(0, length):
        child = breed(pool[i], pool[len(matingpool)-i-1])
        children.append(child)
    return children

def mutatePopulation(constraints, population, mutationRate):
    mutatedPop = []
    
    for ind in population:
        mutatedInd = mutate(constraints, ind, mutationRate)
        mutatedPop.append(mutatedInd)
    return mutatedPop

def nextGeneration(constraints, currentGen, eliteSize, mutationRate, progress):
    popRanked = rankPop(constraints, currentGen)
    progress.append(popRanked[0][1])
    selectionResults = selection(popRanked, eliteSize)
    matingpool = matingPool(currentGen, selectionResults)
    children = breedPopulation(matingpool, eliteSize)
    nextGeneration = mutatePopulation(constraints, children, mutationRate)
    return nextGeneration, progress

def geneticAlgorithm(constraints, popSize, eliteSize, mutationRate, generations, withPlot):
    pop = initialPopulation(popSize, constraints)
    progress = []
    print("Best initial arrangement: " + str(rankPop(constraints, pop)[0][1]))
    for i in range(0, generations):
        pop , progress = nextGeneration(constraints, pop, eliteSize, mutationRate, progress)

    popRanked = rankPop(constraints, pop)
    bestPlanIndex = popRanked[0][0]
    bestPlan = pop[bestPlanIndex]
    print("Best arrangement: " + str(popRanked[0][1]))
    for machines in bestPlan:
        print("===================")
        for machine in machines:
            print(machine.getName())

    if withPlot:
        plt.figure()
        plt.plot(progress)
        plt.ylabel('Function')
        plt.xlabel('Generation')
        plt.show()
    return bestPlan

def printArrangement(constraints, bestPlan, withPlot):
    calculator = PlanCalculator(constraints, bestPlan)
    calculator.calculate()
    greedyArrangement = calculator.getBestArrangement()
    maxTime = 1
    for index, component in enumerate(constraints[componentsKey]):
        print("Cycle time for component {}: {}".format(component.getName(), str(calculator.getCycleTime(index)) ))
        maxTime = max(maxTime, calculator.getCycleTime(index))
    components = {}
    for machineName, machinePlans in greedyArrangement.items():
        print("Machine " + machineName)
        machineNumber = int(machineName[1:])
        for p in machinePlans:
            if p[componentNameKey] not in components:
                components[p[componentNameKey]] = [None] * 10 * math.ceil(maxTime)
            print("  Component : {}".format(p[componentNameKey]))
            print("      time = [{} {}]".format(str(p[timeKey][0]), str(p[timeKey][1])))
            for i in np.arange(p[timeKey][0],p[timeKey][1],0.1):
                components[p[componentNameKey]][int(i*10)] = machineNumber
    if withPlot:
        plt.figure()
        labels = []
        for key, arrayVals in components.items():
            labels.append(key)
            plt.plot(arrayVals)
        plt.ylabel('Machine')
        plt.xlabel('Time')
        plt.legend(labels, loc='upper left')
        plt.show()

constraintsModel = {
    machinesKey:[Machine(name='M1',cost=1.5,processNames=['P1','P2']),Machine(name='M2',cost=0.75,processNames=['P1']),Machine(name='M3',cost=3,processNames=['P1','P2','P3']),
                 Machine(name='M4',cost=2,processNames=['P1','P3']),Machine(name='M5',cost=3,processNames=['P5']),Machine(name='M6',cost=1.5,processNames=['P2','P3'])],
    componentsKey:[
        Component(name='C5', price=25.0, processes=[{processNameKey:'P1', processTimeKey:4.0},{processNameKey:'P2', processTimeKey:2.5},{processNameKey:'P3', processTimeKey:1}]),
        Component(name='C6', price=20.0, processes=[{processNameKey:'P1', processTimeKey:2.0},{processNameKey:'P2', processTimeKey:2},{processNameKey:'P3', processTimeKey:2}]),
        Component(name='C1', price=10.0, processes=[{processNameKey:'P1', processTimeKey:3}]),
        Component(name='C2', price=12.0, processes=[{processNameKey:'P1', processTimeKey:2},{processNameKey:'P2', processTimeKey:3}]),
        Component(name='C3', price=8.0, processes=[{processNameKey:'P2', processTimeKey:3}]),
        Component(name='C4', price=14.0, processes=[{processNameKey:'P2', processTimeKey:3.5},{processNameKey:'P3', processTimeKey:2}]),
        Component(name='C8', price=14.0, processes=[{processNameKey:'P3', processTimeKey:3.5},{processNameKey:'P2', processTimeKey:2}]),
        Component(name='C7', price=15.0, processes=[{processNameKey:'P5', processTimeKey:2.0}])
    ]
}
# score balancer between profitability and product coverage
c1 = 0.6
c2 = 0.4

bestPlan = geneticAlgorithm(constraints=constraintsModel, popSize=100, eliteSize=20, mutationRate=0.1, generations=15, withPlot=True)

printArrangement(constraintsModel, bestPlan, True)