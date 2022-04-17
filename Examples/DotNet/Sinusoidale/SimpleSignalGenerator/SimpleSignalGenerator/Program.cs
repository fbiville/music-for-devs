using NAudio.Wave;
using NAudio.Wave.SampleProviders;
using System;
using System.Threading;

namespace SimpleSignalGenerator
{
    internal class Program
    {
        static void Main(string[] args)
        {
            var waveForm = SignalGeneratorType.Sin;
            var frequency = 440;
            if (args.Length > 0)
            {
                var formParam = args[0];
                switch (formParam.ToLower())
                {
                    case "square": waveForm = SignalGeneratorType.Square; break;
                    case "sin": waveForm = SignalGeneratorType.Sin;break;
                    case "triangle": waveForm = SignalGeneratorType.Triangle; break;
                    default: waveForm = SignalGeneratorType.Sin;break;
                };
            }
            if (args.Length > 1)
            {
                var frequencyParam = args[1];
                int.TryParse(frequencyParam, out frequency);
            }
            TestSignalGen(waveForm, frequency);
            Console.ReadLine();
        }

        private static void TestSignalGen(SignalGeneratorType waveForm, int frequency)
        {
            var A440Sin = new SignalGenerator()
            {
                Gain = 0.2,
                Frequency = frequency,
                Type = waveForm
            }.Take(TimeSpan.FromSeconds(5));


            using var wo1 = new WaveOutEvent();
            wo1.Init(A440Sin);
            wo1.Play();
            while (wo1.PlaybackState == PlaybackState.Playing)
            {
                Thread.Sleep(1000);
            }

        }
    }
}
