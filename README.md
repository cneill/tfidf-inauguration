# Inaugural speech analyzer

## About

This computes the [tf-idf](http://www.tfidf.com/) on presidential inauguration
speeches. The speech texts are included with the repo.

This dataset only covers presidents Ronald Reagan, George H.W. Bush,
Bill Clinton, George W. Bush, Barack Obama, and Donald Trump

The speeches were retrived [here](http://avalon.law.yale.edu/subject_menus/inaug.asp)
and [here](http://www.presidency.ucsb.edu/inaugurals.php).

## Using the tool

```
go build
./inauguration
```

## Data

```
Speech: speeches/barack-obama-1.txt; Words: 2391
TFIDF:
	1. carried: 0.002889
	2. icy: 0.001926
	3. waters: 0.001926
	4. calls: 0.001926
	5. faced: 0.001926
	6. virtue: 0.001926
	7. lower: 0.001926
	8. charter: 0.001926
	9. false: 0.001926
	10. father: 0.001378

Speech: speeches/barack-obama-2.txt; Words: 2090
TFIDF:
	1. requires: 0.003850
	2. complete: 0.003850
	3. she: 0.003080
	4. until: 0.002628
	5. knowing: 0.002310
	6. thrives: 0.002203
	7. train: 0.002203
	8. harm: 0.002203
	9. self-evident: 0.002203
	10. generation's: 0.002203

Speech: speeches/bill-clinton-1.txt; Words: 1583
TFIDF:
	1. season: 0.005818
	2. renewal: 0.004067
	3. whom: 0.004067
	4. change: 0.003503
	5. spring: 0.003050
	6. raised: 0.003050
	7. sake: 0.003050
	8. posterity: 0.003050
	9. serving: 0.002909
	10. compete: 0.002909

Speech: speeches/bill-clinton-2.txt; Words: 2157
TFIDF:
	1. enough: 0.003056
	2. 20th: 0.002238
	3. 21st: 0.002238
	4. waste: 0.002135
	5. saved: 0.002135
	6. information: 0.002135
	7. moved: 0.002135
	8. awful: 0.002135
	9. doors: 0.002135
	10. labors: 0.002135

Speech: speeches/donald-trump-1.txt; Words: 1439
TFIDF:
	1. protected: 0.005592
	2. we've: 0.004800
	3. obama: 0.004800
	4. countries: 0.004800
	5. industry: 0.003200
	6. mountain: 0.003200
	7. exists: 0.003200
	8. transferring: 0.003200
	9. breath: 0.003200
	10. stops: 0.003200

Speech: speeches/george-hw-bush.txt; Words: 2317
TFIDF:
	1. word: 0.004969
	2. breeze: 0.004969
	3. door: 0.003975
	4. fact: 0.003473
	5. other's: 0.002981
	6. blowing: 0.002981
	7. loyal: 0.002981
	8. expression: 0.002981
	9. don't: 0.002845
	10. friends: 0.002393

Speech: speeches/george-w-bush-1.txt; Words: 1579
TFIDF:
	1. story: 0.006262
	2. civility: 0.005833
	3. affirm: 0.003058
	4. directs: 0.002917
	5. stakes: 0.002917
	6. rides: 0.002917
	7. angel: 0.002917
	8. laws: 0.002917
	9. grand: 0.002917
	10. whirlwind: 0.002917

Speech: speeches/george-w-bush-2.txt; Words: 2054
TFIDF:
	1. excuse: 0.003363
	2. fire: 0.003134
	3. defended: 0.002351
	4. ideal: 0.002351
	5. goal: 0.002351
	6. questions: 0.002351
	7. permanent: 0.002351
	8. institutions: 0.002351
	9. accepted: 0.002242
	10. resentment: 0.002242

Speech: speeches/ronald-reagan-1.txt; Words: 2438
TFIDF:
	1. productivity: 0.002833
	2. having: 0.002833
	3. maintaining: 0.002833
	4. burden: 0.002833
	5. weapon: 0.002833
	6. intention: 0.002833
	7. special: 0.002641
	8. heroes: 0.002253
	9. why: 0.001980
	10. front: 0.001980

Speech: speeches/ronald-reagan-2.txt; Words: 2563
TFIDF:
	1. increase: 0.003594
	2. song: 0.002695
	3. government's: 0.002695
	4. weapons: 0.002572
	5. nuclear: 0.002572
	6. human: 0.002434
	7. federal: 0.002143
	8. reduce: 0.002143
	9. spend: 0.001884
	10. number: 0.001884
```
